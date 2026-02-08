package repository

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/grovecj/warzone-stats-tracker/internal/model"
)

type MatchRepo struct {
	pool *pgxpool.Pool
}

func NewMatchRepo(pool *pgxpool.Pool) *MatchRepo {
	return &MatchRepo{pool: pool}
}

func (r *MatchRepo) UpsertBatch(ctx context.Context, playerID string, matches []model.Match) error {
	for _, m := range matches {
		rawJSON, err := json.Marshal(m)
		if err != nil {
			return err
		}
		_, err = r.pool.Exec(ctx, `
			INSERT INTO matches (match_id, player_id, mode, map_name, placement, kills, deaths,
				damage_dealt, damage_taken, gulag_result, match_time, raw_data)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
			ON CONFLICT (match_id, player_id) DO NOTHING
		`, m.MatchID, playerID, m.Mode, m.MapName, m.Placement,
			m.Kills, m.Deaths, m.DamageDealt, m.DamageTaken,
			m.GulagResult, m.MatchTime, rawJSON)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *MatchRepo) GetByPlayerID(ctx context.Context, playerID string, limit, offset int) ([]model.Match, error) {
	if limit <= 0 {
		limit = 20
	}

	rows, err := r.pool.Query(ctx, `
		SELECT id, match_id, player_id, mode, map_name, placement, kills, deaths,
			damage_dealt, damage_taken, gulag_result, match_time, created_at
		FROM matches
		WHERE player_id = $1
		ORDER BY match_time DESC NULLS LAST
		LIMIT $2 OFFSET $3
	`, playerID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []model.Match
	for rows.Next() {
		var m model.Match
		if err := rows.Scan(&m.ID, &m.MatchID, &m.PlayerID, &m.Mode, &m.MapName,
			&m.Placement, &m.Kills, &m.Deaths, &m.DamageDealt, &m.DamageTaken,
			&m.GulagResult, &m.MatchTime, &m.CreatedAt); err != nil {
			return nil, err
		}
		matches = append(matches, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return matches, nil
}

func (r *MatchRepo) CountByPlayerID(ctx context.Context, playerID string) (int, error) {
	var count int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM matches WHERE player_id = $1`, playerID).Scan(&count)
	return count, err
}
