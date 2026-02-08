package service

import (
	"context"
	"log/slog"

	"github.com/grovecj/warzone-stats-tracker/internal/codclient"
	"github.com/grovecj/warzone-stats-tracker/internal/model"
	"github.com/grovecj/warzone-stats-tracker/internal/repository"
)

// MatchListResult contains paginated match data.
type MatchListResult struct {
	Matches []model.Match `json:"matches"`
	Total   int           `json:"total"`
	Limit   int           `json:"limit"`
	Offset  int           `json:"offset"`
}

// MatchService handles match-related business logic.
type MatchService struct {
	codClient  codclient.CodClient
	matchRepo  *repository.MatchRepo
	playerRepo *repository.PlayerRepo
}

// NewMatchService creates a new MatchService.
func NewMatchService(codClient codclient.CodClient, matchRepo *repository.MatchRepo, playerRepo *repository.PlayerRepo) *MatchService {
	return &MatchService{codClient: codClient, matchRepo: matchRepo, playerRepo: playerRepo}
}

// GetRecentMatches fetches matches from the CoD API, persists them, and returns paginated results.
func (s *MatchService) GetRecentMatches(ctx context.Context, platform, gamertag, title, mode string, limit, offset int) (*MatchListResult, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	// Ensure player exists in DB
	player, err := s.playerRepo.GetByPlatformAndTag(ctx, platform, gamertag)
	if err != nil {
		return nil, err
	}
	if player == nil {
		player, err = s.playerRepo.Upsert(ctx, platform, gamertag)
		if err != nil {
			return nil, err
		}
	}

	// Fetch from CoD API
	apiMatches, err := s.codClient.GetRecentMatches(ctx, platform, gamertag, title, mode)
	if err != nil {
		slog.Warn("failed to fetch matches from API, falling back to DB", "error", err)
	} else {
		// Convert and persist
		modelMatches := make([]model.Match, 0, len(apiMatches))
		for _, m := range apiMatches {
			modelMatches = append(modelMatches, model.Match{
				MatchID:     m.MatchID,
				PlayerID:    player.ID,
				Mode:        m.Mode,
				MapName:     m.Map,
				Placement:   m.Placement,
				Kills:       m.Kills,
				Deaths:      m.Deaths,
				DamageDealt: m.DamageDealt,
				DamageTaken: m.DamageTaken,
				GulagResult: m.GulagResult,
				MatchTime:   m.MatchTime,
			})
		}
		if upsertErr := s.matchRepo.UpsertBatch(ctx, player.ID, modelMatches); upsertErr != nil {
			slog.Warn("failed to persist matches", "error", upsertErr)
		}
	}

	// Read from DB with pagination
	matches, err := s.matchRepo.GetByPlayerID(ctx, player.ID, limit, offset)
	if err != nil {
		return nil, err
	}

	total, err := s.matchRepo.CountByPlayerID(ctx, player.ID)
	if err != nil {
		slog.Warn("failed to count matches", "error", err)
		total = len(matches)
	}

	return &MatchListResult{
		Matches: matches,
		Total:   total,
		Limit:   limit,
		Offset:  offset,
	}, nil
}
