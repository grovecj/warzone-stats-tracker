package codclient

import "time"

// PlayerStats represents lifetime player statistics from the CoD API.
type PlayerStats struct {
	Platform    string  `json:"platform"`
	Gamertag    string  `json:"gamertag"`
	Level       int     `json:"level"`
	Prestige    int     `json:"prestige"`
	Kills       int     `json:"kills"`
	Deaths      int     `json:"deaths"`
	KDRatio     float64 `json:"kdRatio"`
	Wins        int     `json:"wins"`
	Losses      int     `json:"losses"`
	WinPct      float64 `json:"winPct"`
	ScorePerMin float64 `json:"scorePerMin"`
	Headshots   int     `json:"headshots"`
	TimePlayed  int     `json:"timePlayed"`
	MatchesPlayed int   `json:"matchesPlayed"`
}

// Match represents a single match from the CoD API.
type Match struct {
	MatchID     string    `json:"matchID"`
	Mode        string    `json:"mode"`
	Map         string    `json:"map"`
	Placement   int       `json:"placement"`
	Kills       int       `json:"kills"`
	Deaths      int       `json:"deaths"`
	KDRatio     float64   `json:"kdRatio"`
	DamageDealt int       `json:"damageDealt"`
	DamageTaken int       `json:"damageTaken"`
	GulagResult string    `json:"gulagResult"`
	MatchTime   time.Time `json:"matchTime"`
}
