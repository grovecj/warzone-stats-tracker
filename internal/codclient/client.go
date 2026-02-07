package codclient

import "context"

// CodClient defines the interface for interacting with the Call of Duty API.
// Full implementation in issue #6.
type CodClient interface {
	GetPlayerStats(ctx context.Context, platform, gamertag, mode string) (*PlayerStats, error)
	GetRecentMatches(ctx context.Context, platform, gamertag string) ([]Match, error)
}
