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
  modeBreakdown?: Record<string, ModeStats>
}

export interface ModeStats {
  kills: number
  deaths: number
  kdRatio: number
  wins: number
  losses: number
  matchesPlayed: number
  scorePerMin: number
  timePlayed: number
  topFive: number
  topTen: number
  topTwentyFive: number
}

export interface Player {
  id: string
  platform: string
  gamertag: string
  activisionId?: string
  lastFetchedAt?: string
  createdAt: string
  updatedAt: string
}

export interface PlayerSearchResult {
  playerId: string
  platform: string
  gamertag: string
  stats: PlayerStats
}

export interface Match {
  id?: string
  matchId?: string
  matchID?: string
  playerId?: string
  mode: string
  map?: string
  mapName?: string
  placement: number
  kills: number
  deaths: number
  kdRatio?: number
  damageDealt: number
  damageTaken: number
  gulagResult: string
  duration?: number
  matchTime: string
  createdAt?: string
}

export interface MatchListResult {
  matches: Match[]
  total: number
  limit: number
  offset: number
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
