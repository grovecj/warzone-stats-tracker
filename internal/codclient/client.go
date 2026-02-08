package codclient

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"sync"
	"time"

	"resty.dev/v3"
)

const (
	defaultTitle = "mw" // Modern Warfare / Warzone title code
)

// CodClient defines the interface for interacting with the Call of Duty API.
type CodClient interface {
	GetPlayerStats(ctx context.Context, platform, gamertag, mode string) (*PlayerStats, error)
	GetRecentMatches(ctx context.Context, platform, gamertag string) ([]Match, error)
	UpdateToken(newToken string)
}

type client struct {
	http    *resty.Client
	baseURL string
	mu      sync.RWMutex
	token   string
}

// New creates a new CoD API client.
func New(baseURL, ssoToken string) CodClient {
	c := resty.New()
	c.SetBaseURL(baseURL)
	c.SetTimeout(10 * time.Second)
	c.SetRetryCount(3)
	c.SetRetryWaitTime(1 * time.Second)
	c.SetRetryMaxWaitTime(5 * time.Second)
	c.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	c.SetHeader("Accept", "application/json")

	return &client{http: c, baseURL: baseURL, token: ssoToken}
}

func (c *client) authCookie() *http.Cookie {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return &http.Cookie{Name: "ACT_SSO_COOKIE", Value: c.token}
}

func (c *client) GetPlayerStats(ctx context.Context, platform, gamertag, mode string) (*PlayerStats, error) {
	if mode == "" {
		mode = "wz"
	}

	encodedTag := url.PathEscape(gamertag)
	endpoint := fmt.Sprintf("/stats/cod/v1/title/%s/platform/%s/gamer/%s/profile/type/%s",
		defaultTitle, platform, encodedTag, mode)

	resp, err := c.http.R().
		SetContext(ctx).
		SetCookie(c.authCookie()).
		Get(endpoint)
	if err != nil {
		slog.Error("cod api request failed", "endpoint", endpoint, "error", err)
		return nil, ErrAPIUnavailable
	}

	if err := c.checkResponse(resp); err != nil {
		return nil, err
	}

	var profileResp profileResponse
	if err := json.Unmarshal(resp.Bytes(), &profileResp); err != nil {
		return nil, fmt.Errorf("decoding profile response: %w", err)
	}

	stats := c.mapProfileToStats(profileResp, platform, gamertag)
	return stats, nil
}

func (c *client) GetRecentMatches(ctx context.Context, platform, gamertag string) ([]Match, error) {
	encodedTag := url.PathEscape(gamertag)
	endpoint := fmt.Sprintf("/crm/cod/v2/title/%s/platform/%s/gamer/%s/matches/wz/start/0/end/0/details",
		defaultTitle, platform, encodedTag)

	resp, err := c.http.R().
		SetContext(ctx).
		SetCookie(c.authCookie()).
		Get(endpoint)
	if err != nil {
		slog.Error("cod api request failed", "endpoint", endpoint, "error", err)
		return nil, ErrAPIUnavailable
	}

	if err := c.checkResponse(resp); err != nil {
		return nil, err
	}

	var matchResp matchesResponse
	if err := json.Unmarshal(resp.Bytes(), &matchResp); err != nil {
		return nil, fmt.Errorf("decoding matches response: %w", err)
	}

	matches := make([]Match, 0, len(matchResp.Data.Matches))
	for _, m := range matchResp.Data.Matches {
		gulag := ""
		if m.PlayerStats.GulagKills > 0 {
			gulag = "win"
		} else if m.PlayerStats.GulagDeaths > 0 {
			gulag = "loss"
		}

		matches = append(matches, Match{
			MatchID:     m.MatchID,
			Mode:        m.Mode,
			Map:         m.Map,
			Placement:   m.PlayerStats.TeamPlacement,
			Kills:       m.PlayerStats.Kills,
			Deaths:      m.PlayerStats.Deaths,
			KDRatio:     m.PlayerStats.KDRatio,
			DamageDealt: m.PlayerStats.DamageDone,
			DamageTaken: m.PlayerStats.DamageTaken,
			GulagResult: gulag,
			Duration:    m.Duration,
			MatchTime:   time.Unix(int64(m.UTCStartSeconds), 0),
		})
	}

	return matches, nil
}

// UpdateToken replaces the SSO token at runtime (for admin token refresh).
func (c *client) UpdateToken(newToken string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.token = newToken
	slog.Info("cod api sso token updated")
}

func (c *client) checkResponse(resp *resty.Response) error {
	switch resp.StatusCode() {
	case http.StatusOK:
		return nil
	case http.StatusUnauthorized:
		return ErrTokenExpired
	case http.StatusForbidden:
		return ErrPrivateProfile
	case http.StatusNotFound:
		return ErrPlayerNotFound
	case http.StatusTooManyRequests:
		return ErrRateLimited
	default:
		if resp.StatusCode() >= 500 {
			return ErrAPIUnavailable
		}
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode(), resp.String())
	}
}

func (c *client) mapProfileToStats(resp profileResponse, platform, gamertag string) *PlayerStats {
	stats := &PlayerStats{
		Platform: platform,
		Gamertag: gamertag,
		Level:    resp.Data.Level,
		Prestige: resp.Data.Prestige,
	}

	if props, ok := resp.Data.Lifetime.All["properties"]; ok {
		stats.Kills = int(props.Kills)
		stats.Deaths = int(props.Deaths)
		stats.KDRatio = props.KDRatio
		stats.Wins = int(props.Wins)
		stats.Losses = int(props.Losses)
		stats.WinPct = props.WinPct
		stats.ScorePerMin = props.ScorePerMin
		stats.Headshots = int(props.Headshots)
		stats.TimePlayed = int(props.TimePlayed)
		stats.MatchesPlayed = int(props.MatchesPlayed)
		stats.TopFive = int(props.TopFive)
		stats.TopTen = int(props.TopTen)
		stats.TopTwentyFive = int(props.TopTwentyFive)
		stats.Assists = int(props.Assists)
		stats.DamageDone = int(props.DamageDone)
	}

	return stats
}
