package service

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/grovecj/warzone-stats-tracker/internal/codclient"
	"github.com/grovecj/warzone-stats-tracker/internal/repository"
)

// PlayerSearchResult contains the player info and a stats summary from a search.
type PlayerSearchResult struct {
	PlayerID string                 `json:"playerId"`
	Platform string                 `json:"platform"`
	Gamertag string                 `json:"gamertag"`
	Stats    *codclient.PlayerStats `json:"stats"`
}

// PlayerService handles player-related business logic.
type PlayerService struct {
	codClient  codclient.CodClient
	playerRepo *repository.PlayerRepo
}

// NewPlayerService creates a new PlayerService.
func NewPlayerService(codClient codclient.CodClient, playerRepo *repository.PlayerRepo) *PlayerService {
	return &PlayerService{codClient: codClient, playerRepo: playerRepo}
}

// SearchPlayer verifies a player exists via the CoD API, persists them, and returns search results.
func (s *PlayerService) SearchPlayer(ctx context.Context, platform, gamertag, title, mode string) (*PlayerSearchResult, error) {
	if title == "" {
		title = "mw"
	}
	if mode == "" {
		mode = "wz"
	}

	stats, err := s.codClient.GetPlayerStats(ctx, platform, gamertag, title, mode)
	if err != nil {
		return nil, err
	}

	player, err := s.playerRepo.Upsert(ctx, platform, gamertag)
	if err != nil {
		slog.Error("failed to upsert player", "platform", platform, "gamertag", gamertag, "error", err)
		return nil, err
	}

	statsJSON, err := json.Marshal(stats)
	if err == nil {
		if saveErr := s.playerRepo.SaveStatsSnapshot(ctx, player.ID, mode, statsJSON); saveErr != nil {
			slog.Warn("failed to save stats snapshot", "player_id", player.ID, "error", saveErr)
		}
	}

	return &PlayerSearchResult{
		PlayerID: player.ID,
		Platform: platform,
		Gamertag: gamertag,
		Stats:    stats,
	}, nil
}

// GetPlayerStats fetches player stats from the CoD API, upserts the player, and saves a snapshot.
func (s *PlayerService) GetPlayerStats(ctx context.Context, platform, gamertag, title, mode string) (*codclient.PlayerStats, error) {
	if title == "" {
		title = "mw"
	}
	if mode == "" {
		mode = "wz"
	}

	stats, err := s.codClient.GetPlayerStats(ctx, platform, gamertag, title, mode)
	if err != nil {
		return nil, err
	}

	player, err := s.playerRepo.Upsert(ctx, platform, gamertag)
	if err != nil {
		slog.Error("failed to upsert player", "platform", platform, "gamertag", gamertag, "error", err)
		return stats, nil // return stats even if DB write fails
	}

	statsJSON, err := json.Marshal(stats)
	if err == nil {
		if saveErr := s.playerRepo.SaveStatsSnapshot(ctx, player.ID, mode, statsJSON); saveErr != nil {
			slog.Warn("failed to save stats snapshot", "player_id", player.ID, "error", saveErr)
		}
	}

	return stats, nil
}
