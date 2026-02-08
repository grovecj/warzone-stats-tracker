export interface PlayerStats {
  platform: string
  gamertag: string
  level: number
  prestige: number
  kills: number
  deaths: number
  kdRatio: number
  wins: number
  losses: number
  winPct: number
  scorePerMin: number
  headshots: number
  timePlayed: number
  matchesPlayed: number
  topFive: number
  topTen: number
  topTwentyFive: number
  assists: number
  damageDone: number
}

export interface Match {
  matchID: string
  mode: string
  map: string
  placement: number
  kills: number
  deaths: number
  kdRatio: number
  damageDealt: number
  damageTaken: number
  gulagResult: string
  duration: number
  matchTime: string
}

export interface Squad {
  id: string
  name: string
  members: SquadMember[]
  createdAt: string
  updatedAt: string
}

export interface SquadMember {
  id: string
  platform: string
  gamertag: string
}
