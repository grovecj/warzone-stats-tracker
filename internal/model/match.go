package model

import "time"

type Match struct {
	ID          string    `json:"id"`
	MatchID     string    `json:"matchId"`
	PlayerID    string    `json:"playerId"`
	Mode        string    `json:"mode"`
	MapName     string    `json:"mapName"`
	Placement   int       `json:"placement"`
	Kills       int       `json:"kills"`
	Deaths      int       `json:"deaths"`
	DamageDealt int       `json:"damageDealt"`
	DamageTaken int       `json:"damageTaken"`
	GulagResult string    `json:"gulagResult,omitempty"`
	MatchTime   time.Time `json:"matchTime"`
	CreatedAt   time.Time `json:"createdAt"`
}
