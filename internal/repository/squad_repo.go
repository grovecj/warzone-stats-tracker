package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/grovecj/warzone-stats-tracker/internal/model"
)

type SquadRepo struct {
	pool *pgxpool.Pool
}

func NewSquadRepo(pool *pgxpool.Pool) *SquadRepo {
	return &SquadRepo{pool: pool}
}

func (r *SquadRepo) Create(ctx context.Context, name string) (*model.Squad, error) {
	var s model.Squad
	err := r.pool.QueryRow(ctx, `
		INSERT INTO squads (name) VALUES ($1)
		RETURNING id, name, created_at, updated_at
	`, name).Scan(&s.ID, &s.Name, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *SquadRepo) GetByID(ctx context.Context, id string) (*model.Squad, error) {
	var s model.Squad
	err := r.pool.QueryRow(ctx, `
		SELECT id, name, created_at, updated_at FROM squads WHERE id = $1
	`, id).Scan(&s.ID, &s.Name, &s.CreatedAt, &s.UpdatedAt)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, `
		SELECT p.id, p.platform, p.gamertag, p.activision_id, p.last_fetched_at, p.created_at, p.updated_at
		FROM players p
		JOIN squad_members sm ON sm.player_id = p.id
		WHERE sm.squad_id = $1
		ORDER BY sm.added_at
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Player
		if err := rows.Scan(&p.ID, &p.Platform, &p.Gamertag, &p.ActivisionID,
			&p.LastFetchedAt, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		s.Members = append(s.Members, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *SquadRepo) Update(ctx context.Context, id, name string) error {
	_, err := r.pool.Exec(ctx, `
		UPDATE squads SET name = $2, updated_at = NOW() WHERE id = $1
	`, id, name)
	return err
}

func (r *SquadRepo) Delete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM squads WHERE id = $1`, id)
	return err
}

func (r *SquadRepo) AddMember(ctx context.Context, squadID, playerID string) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO squad_members (squad_id, player_id) VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`, squadID, playerID)
	return err
}

func (r *SquadRepo) RemoveMember(ctx context.Context, squadID, playerID string) error {
	_, err := r.pool.Exec(ctx, `
		DELETE FROM squad_members WHERE squad_id = $1 AND player_id = $2
	`, squadID, playerID)
	return err
}

func (r *SquadRepo) MemberCount(ctx context.Context, squadID string) (int, error) {
	var count int
	err := r.pool.QueryRow(ctx, `
		SELECT COUNT(*) FROM squad_members WHERE squad_id = $1
	`, squadID).Scan(&count)
	return count, err
}
