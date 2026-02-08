package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/grovecj/warzone-stats-tracker/internal/model"
)

type PlayerRepo struct {
	pool *pgxpool.Pool
}

func NewPlayerRepo(pool *pgxpool.Pool) *PlayerRepo {
	return &PlayerRepo{pool: pool}
}

func (r *PlayerRepo) Upsert(ctx context.Context, platform, gamertag string) (*model.Player, error) {
	var p model.Player
	err := r.pool.QueryRow(ctx, `
		INSERT INTO players (platform, gamertag, last_fetched_at)
		VALUES ($1, $2, NOW())
		ON CONFLICT (platform, gamertag)
		DO UPDATE SET last_fetched_at = NOW(), updated_at = NOW()
		RETURNING id, platform, gamertag, activision_id, last_fetched_at, created_at, updated_at
	`, platform, gamertag).Scan(
		&p.ID, &p.Platform, &p.Gamertag, &p.ActivisionID,
		&p.LastFetchedAt, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PlayerRepo) GetByPlatformAndTag(ctx context.Context, platform, gamertag string) (*model.Player, error) {
	var p model.Player
	err := r.pool.QueryRow(ctx, `
		SELECT id, platform, gamertag, activision_id, last_fetched_at, created_at, updated_at
		FROM players WHERE platform = $1 AND gamertag = $2
	`, platform, gamertag).Scan(
		&p.ID, &p.Platform, &p.Gamertag, &p.ActivisionID,
		&p.LastFetchedAt, &p.CreatedAt, &p.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PlayerRepo) GetByID(ctx context.Context, id string) (*model.Player, error) {
	var p model.Player
	err := r.pool.QueryRow(ctx, `
		SELECT id, platform, gamertag, activision_id, last_fetched_at, created_at, updated_at
		FROM players WHERE id = $1
	`, id).Scan(
		&p.ID, &p.Platform, &p.Gamertag, &p.ActivisionID,
		&p.LastFetchedAt, &p.CreatedAt, &p.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PlayerRepo) SaveStatsSnapshot(ctx context.Context, playerID, mode string, statsData any) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO player_stats (player_id, mode, stats_data) VALUES ($1, $2, $3)
	`, playerID, mode, statsData)
	return err
}

func (r *PlayerRepo) GetLatestStats(ctx context.Context, playerID, mode string) (any, *time.Time, error) {
	var statsData any
	var fetchedAt time.Time
	err := r.pool.QueryRow(ctx, `
		SELECT stats_data, fetched_at FROM player_stats
		WHERE player_id = $1 AND mode = $2
		ORDER BY fetched_at DESC LIMIT 1
	`, playerID, mode).Scan(&statsData, &fetchedAt)
	if err == pgx.ErrNoRows {
		return nil, nil, nil
	}
	if err != nil {
		return nil, nil, err
	}
	return statsData, &fetchedAt, nil
}
