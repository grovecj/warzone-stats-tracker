package codclient

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"resty.dev/v3"
)

// CodClient defines the interface for interacting with the Call of Duty API.
type CodClient interface {
	GetPlayerStats(ctx context.Context, platform, gamertag, title, mode string) (*PlayerStats, error)
	GetRecentMatches(ctx context.Context, platform, gamertag, title, mode string) ([]Match, error)
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
	c.SetRedirectPolicy(resty.NoRedirectPolicy())

	return &client{http: c, baseURL: baseURL, token: ssoToken}
}

func (c *client) authCookie() *http.Cookie {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return &http.Cookie{Name: "ACT_SSO_COOKIE", Value: c.token}
}

func (c *client) GetPlayerStats(ctx context.Context, platform, gamertag, title, mode string) (*PlayerStats, error) {
	if title == "" {
		title = "mw"
	}
	if mode == "" {
		mode = "wz"
	}

	encodedTag := url.PathEscape(gamertag)
	endpoint := fmt.Sprintf("/stats/cod/v1/title/%s/platform/%s/gamer/%s/profile/type/%s",
		title, platform, encodedTag, mode)

	resp, err := c.doRequest(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	var profileResp profileResponse
	if err := json.Unmarshal(resp.Bytes(), &profileResp); err != nil {
		return nil, fmt.Errorf("decoding profile response: %w", err)
	}

	stats := c.mapProfileToStats(profileResp, platform, gamertag)
	return stats, nil
}

func (c *client) GetRecentMatches(ctx context.Context, platform, gamertag, title, mode string) ([]Match, error) {
	if title == "" {
		title = "mw"
	}
	if mode == "" {
		mode = "wz"
	}

	encodedTag := url.PathEscape(gamertag)
	endpoint := fmt.Sprintf("/crm/cod/v2/title/%s/platform/%s/gamer/%s/matches/%s/start/0/end/0/details",
		title, platform, encodedTag, mode)

	resp, err := c.doRequest(ctx, endpoint)
	if err != nil {
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

// doRequest performs a GET request, handling redirect errors as expired tokens.
func (c *client) doRequest(ctx context.Context, endpoint string) (*resty.Response, error) {
	resp, err := c.http.R().
		SetContext(ctx).
		SetCookie(c.authCookie()).
		Get(endpoint)
	if err != nil {
		// resty returns an error on redirect when NoRedirectPolicy is set,
		// but the response is still populated
		if resp != nil && resp.StatusCode() >= 300 && resp.StatusCode() < 400 {
			slog.Warn("cod api redirected, token likely expired",
				"status", resp.StatusCode(),
				"location", resp.Header().Get("Location"))
			return nil, ErrTokenExpired
		}
		slog.Error("cod api request failed", "endpoint", endpoint, "error", err)
		return nil, ErrAPIUnavailable
	}

	if err := c.checkResponse(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *client) checkResponse(resp *resty.Response) error {
	switch resp.StatusCode() {
	case http.StatusOK:
		body := resp.String()
		// CoD API sometimes returns 200 with HTML (login page) instead of JSON
		if len(body) > 0 && body[0] == '<' {
			slog.Error("cod api returned html instead of json",
				"content_type", resp.Header().Get("Content-Type"),
				"body_prefix", body[:min(200, len(body))])
			return ErrTokenExpired
		}
		// CoD API returns 200 with {"status":"error"} for business-logic errors
		var envelope struct {
			Status string `json:"status"`
			Data   struct {
				Message string `json:"message"`
			} `json:"data"`
		}
		if json.Unmarshal([]byte(body), &envelope) == nil && envelope.Status == "error" {
			return c.mapAPIError(envelope.Data.Message)
		}
		return nil
	case http.StatusUnauthorized:
		return ErrTokenExpired
	case http.StatusForbidden:
		return ErrPrivateProfile
	case http.StatusNotFound:
		return ErrPlayerNotFound
	case http.StatusTooManyRequests:
		return ErrRateLimited
	case http.StatusMovedPermanently, http.StatusFound, http.StatusTemporaryRedirect:
		// CoD API redirects to login/store page when token is expired
		slog.Warn("cod api redirected, token likely expired",
			"status", resp.StatusCode(),
			"location", resp.Header().Get("Location"))
		return ErrTokenExpired
	default:
		if resp.StatusCode() >= 500 {
			return ErrAPIUnavailable
		}
		body := resp.String()
		slog.Error("cod api unexpected status", "status", resp.StatusCode(),
			"body_prefix", body[:min(200, len(body))])
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode(), body)
	}
}

// mapAPIError converts CoD API error messages into sentinel errors.
func (c *client) mapAPIError(msg string) error {
	slog.Warn("cod api returned error", "message", msg)
	switch {
	case strings.Contains(msg, "not authenticated"):
		return ErrTokenExpired
	case strings.Contains(msg, "not allowed"):
		return ErrPlayerNotFound
	case strings.Contains(msg, "user not found"):
		return ErrPlayerNotFound
	case strings.Contains(msg, "rate limit"):
		return ErrRateLimited
	default:
		return fmt.Errorf("cod api error: %s", msg)
	}
}

func (c *client) mapProfileToStats(resp profileResponse, platform, gamertag string) *PlayerStats {
	stats := &PlayerStats{
		Platform: platform,
		Gamertag: gamertag,
		Level:    int(resp.Data.Level),
		Prestige: int(resp.Data.Prestige),
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

	// Parse per-mode breakdown from resp.Data.Lifetime.Mode
	stats.ModeBreakdown = c.parseModeBreakdown(resp.Data.Lifetime.Mode)

	return stats
}

// parseModeBreakdown extracts per-mode stats from the API's Mode map.
func (c *client) parseModeBreakdown(modeData map[string]any) map[string]ModeStats {
	if len(modeData) == 0 {
		return nil
	}

	breakdown := make(map[string]ModeStats, len(modeData))
	for modeName, modeVal := range modeData {
		modeMap, ok := modeVal.(map[string]any)
		if !ok {
			continue
		}
		propsVal, ok := modeMap["properties"]
		if !ok {
			continue
		}
		props, ok := propsVal.(map[string]any)
		if !ok {
			continue
		}

		breakdown[modeName] = ModeStats{
			Kills:         toInt(props["kills"]),
			Deaths:        toInt(props["deaths"]),
			KDRatio:       toFloat(props["kdRatio"]),
			Wins:          toInt(props["wins"]),
			Losses:        toInt(props["losses"]),
			MatchesPlayed: toInt(props["matchesPlayed"]),
			ScorePerMin:   toFloat(props["scorePerMinute"]),
			TimePlayed:    toInt(props["timePlayed"]),
			TopFive:       toInt(props["topFive"]),
			TopTen:        toInt(props["topTen"]),
			TopTwentyFive: toInt(props["topTwentyFive"]),
		}
	}

	if len(breakdown) == 0 {
		return nil
	}
	return breakdown
}

func toFloat(v any) float64 {
	switch n := v.(type) {
	case float64:
		return n
	case int:
		return float64(n)
	default:
		return 0
	}
}

func toInt(v any) int {
	switch n := v.(type) {
	case float64:
		return int(n)
	case int:
		return n
	default:
		return 0
	}
}
