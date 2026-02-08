package cache

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/grovecj/warzone-stats-tracker/internal/codclient"
)

type entry struct {
	value     any
	expiresAt time.Time
}

func (e entry) isExpired() bool {
	return time.Now().After(e.expiresAt)
}

// CachedClient wraps a CodClient with in-process TTL caching.
type CachedClient struct {
	inner     codclient.CodClient
	mu        sync.RWMutex
	store     map[string]entry
	statsTTL  time.Duration
	matchTTL  time.Duration
}

// Config holds cache TTL settings.
type Config struct {
	StatsTTL time.Duration
	MatchTTL time.Duration
}

func DefaultConfig() Config {
	return Config{
		StatsTTL: 5 * time.Minute,
		MatchTTL: 2 * time.Minute,
	}
}

// New wraps a CodClient with caching.
func New(inner codclient.CodClient, cfg Config) *CachedClient {
	c := &CachedClient{
		inner:    inner,
		store:    make(map[string]entry),
		statsTTL: cfg.StatsTTL,
		matchTTL: cfg.MatchTTL,
	}
	go c.evictLoop()
	return c
}

func (c *CachedClient) GetPlayerStats(ctx context.Context, platform, gamertag, mode string) (*codclient.PlayerStats, error) {
	key := fmt.Sprintf("stats:%s:%s:%s", platform, gamertag, mode)

	if val, hit := c.get(key); hit {
		slog.Debug("cache hit", "key", key)
		return val.(*codclient.PlayerStats), nil
	}

	stats, err := c.inner.GetPlayerStats(ctx, platform, gamertag, mode)
	if err != nil {
		// Serve stale data if available
		if val, ok := c.getStale(key); ok {
			slog.Warn("serving stale cache due to API error", "key", key, "error", err)
			return val.(*codclient.PlayerStats), nil
		}
		return nil, err
	}

	c.set(key, stats, c.statsTTL)
	return stats, nil
}

func (c *CachedClient) GetRecentMatches(ctx context.Context, platform, gamertag string) ([]codclient.Match, error) {
	key := fmt.Sprintf("matches:%s:%s", platform, gamertag)

	if val, hit := c.get(key); hit {
		slog.Debug("cache hit", "key", key)
		return val.([]codclient.Match), nil
	}

	matches, err := c.inner.GetRecentMatches(ctx, platform, gamertag)
	if err != nil {
		if val, ok := c.getStale(key); ok {
			slog.Warn("serving stale cache due to API error", "key", key, "error", err)
			return val.([]codclient.Match), nil
		}
		return nil, err
	}

	c.set(key, matches, c.matchTTL)
	return matches, nil
}

// UpdateToken passes through to the inner client.
func (c *CachedClient) UpdateToken(newToken string) {
	c.inner.UpdateToken(newToken)
}

// CacheInfo returns hit/miss status and age for use in response headers.
func (c *CachedClient) CacheInfo(key string) (hit bool, ageSeconds int, stale bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	e, exists := c.store[key]
	if !exists {
		return false, 0, false
	}

	age := int(time.Since(e.expiresAt.Add(-c.statsTTL)).Seconds())
	if age < 0 {
		age = 0
	}

	return true, age, e.isExpired()
}

func (c *CachedClient) get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	e, ok := c.store[key]
	if !ok || e.isExpired() {
		return nil, false
	}
	return e.value, true
}

func (c *CachedClient) getStale(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	e, ok := c.store[key]
	if !ok {
		return nil, false
	}
	return e.value, true
}

func (c *CachedClient) set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.store[key] = entry{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
}

func (c *CachedClient) evictLoop() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		// Evict entries that have been stale for more than 1 hour
		cutoff := time.Now().Add(-1 * time.Hour)
		for k, e := range c.store {
			if e.expiresAt.Before(cutoff) {
				delete(c.store, k)
			}
		}
		c.mu.Unlock()
	}
}
