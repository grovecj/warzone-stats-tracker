package model

import "time"

type Player struct {
	ID            string    `json:"id"`
	Platform      string    `json:"platform"`
	Gamertag      string    `json:"gamertag"`
	ActivisionID  *string   `json:"activisionId,omitempty"`
	LastFetchedAt *time.Time `json:"lastFetchedAt,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
