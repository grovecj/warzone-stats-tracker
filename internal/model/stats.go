package model

import "time"

type PlayerStatsSnapshot struct {
	ID        string    `json:"id"`
	PlayerID  string    `json:"playerId"`
	Mode      string    `json:"mode"`
	StatsData any       `json:"statsData"`
	FetchedAt time.Time `json:"fetchedAt"`
}
