package codclient

import "time"

// PlayerStats represents lifetime player statistics from the CoD API.
type PlayerStats struct {
	Platform      string  `json:"platform"`
	Gamertag      string  `json:"gamertag"`
	Level         int     `json:"level"`
	Prestige      int     `json:"prestige"`
	Kills         int     `json:"kills"`
	Deaths        int     `json:"deaths"`
	KDRatio       float64 `json:"kdRatio"`
	Wins          int     `json:"wins"`
	Losses        int     `json:"losses"`
	WinPct        float64 `json:"winPct"`
	ScorePerMin   float64 `json:"scorePerMin"`
	Headshots     int     `json:"headshots"`
	TimePlayed    int     `json:"timePlayed"`
	MatchesPlayed int     `json:"matchesPlayed"`
	TopFive       int     `json:"topFive"`
	TopTen        int     `json:"topTen"`
	TopTwentyFive int     `json:"topTwentyFive"`
	Assists       int     `json:"assists"`
	DamageDone    int     `json:"damageDone"`
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
	Duration    int       `json:"duration"`
	MatchTime   time.Time `json:"matchTime"`
	RawData     any       `json:"rawData,omitempty"`
}

// apiResponse is the wrapper returned by the CoD API.
type apiResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

// profileResponse maps the nested profile endpoint response.
type profileResponse struct {
	Status string `json:"status"`
	Data   struct {
		Type     string `json:"type"`
		Message  string `json:"message"`
		Lifetime struct {
			All  map[string]statsBlock `json:"all"`
			Mode map[string]any        `json:"mode"`
		} `json:"lifetime"`
		Level   int `json:"level"`
		Prestige int `json:"prestige"`
	} `json:"data"`
}

type statsBlock struct {
	Kills         float64 `json:"kills"`
	Deaths        float64 `json:"deaths"`
	KDRatio       float64 `json:"kdRatio"`
	Wins          float64 `json:"wins"`
	Losses        float64 `json:"losses"`
	WinPct        float64 `json:"wlRatio"`
	ScorePerMin   float64 `json:"scorePerMinute"`
	Headshots     float64 `json:"headshots"`
	TimePlayed    float64 `json:"timePlayed"`
	MatchesPlayed float64 `json:"matchesPlayed"`
	TopFive       float64 `json:"topFive"`
	TopTen        float64 `json:"topTen"`
	TopTwentyFive float64 `json:"topTwentyFive"`
	Assists       float64 `json:"assists"`
	DamageDone    float64 `json:"damageDone"`
}

// matchesResponse maps the matches endpoint response.
type matchesResponse struct {
	Status string `json:"status"`
	Data   struct {
		Matches []matchData `json:"matches"`
		Message string      `json:"message"`
	} `json:"data"`
}

type matchData struct {
	MatchID       string  `json:"matchID"`
	Mode          string  `json:"mode"`
	Map           string  `json:"map"`
	PlayerStats   matchPlayerStats `json:"playerStats"`
	Duration      int     `json:"duration"`
	UTCStartSeconds float64 `json:"utcStartSeconds"`
	RawData       any     `json:"-"`
}

type matchPlayerStats struct {
	Kills          int     `json:"kills"`
	Deaths         int     `json:"deaths"`
	KDRatio        float64 `json:"kdRatio"`
	DamageDone     int     `json:"damageDone"`
	DamageTaken    int     `json:"damageTaken"`
	TeamPlacement  int     `json:"teamPlacement"`
	GulagKills     int     `json:"gulagKills"`
	GulagDeaths    int     `json:"gulagDeaths"`
}
