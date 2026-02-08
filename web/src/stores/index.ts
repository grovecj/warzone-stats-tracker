import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { PlayerStats, Match, PlayerSearchResult, MatchListResult } from '@/types'

export const usePlayerStore = defineStore('player', () => {
  const statsLoading = ref(false)
  const matchesLoading = ref(false)
  const statsError = ref<string | null>(null)
  const matchesError = ref<string | null>(null)
  const stats = ref<PlayerStats | null>(null)
  const matches = ref<Match[]>([])
  const matchesTotal = ref(0)

  async function searchPlayer(
    platform: string,
    gamertag: string,
    title = 'mw',
    mode = 'wz',
  ): Promise<PlayerSearchResult> {
    const tag = encodeURIComponent(gamertag)
    const res = await fetch(
      `/api/v1/players/search?platform=${encodeURIComponent(platform)}&gamertag=${tag}&title=${encodeURIComponent(title)}&mode=${encodeURIComponent(mode)}`,
    )
    if (!res.ok) {
      const text = await res.text()
      let message = 'Failed to search player'
      try {
        const body = JSON.parse(text)
        message = body.message || message
      } catch {
        // Response was not JSON
      }
      throw new Error(message)
    }
    return await res.json()
  }

  async function fetchStats(platform: string, gamertag: string, title = 'mw', mode = 'wz') {
    statsLoading.value = true
    statsError.value = null
    try {
      const tag = encodeURIComponent(gamertag)
      const res = await fetch(
        `/api/v1/players/${platform}/${tag}/stats?title=${encodeURIComponent(title)}&mode=${encodeURIComponent(mode)}`,
      )
      if (!res.ok) {
        const text = await res.text()
        let message = 'Failed to fetch stats'
        try {
          const body = JSON.parse(text)
          message = body.message || message
        } catch {
          // Response was not JSON
        }
        throw new Error(message)
      }
      stats.value = await res.json()
    } catch (e: unknown) {
      statsError.value = e instanceof Error ? e.message : 'An unexpected error occurred'
    } finally {
      statsLoading.value = false
    }
  }

  async function fetchMatches(platform: string, gamertag: string, title = 'mw', mode = 'wz', limit = 20, offset = 0) {
    matchesLoading.value = true
    matchesError.value = null
    try {
      const tag = encodeURIComponent(gamertag)
      const res = await fetch(
        `/api/v1/players/${platform}/${tag}/matches?title=${encodeURIComponent(title)}&mode=${encodeURIComponent(mode)}&limit=${limit}&offset=${offset}`,
      )
      if (!res.ok) {
        const text = await res.text()
        let message = 'Failed to fetch matches'
        try {
          const body = JSON.parse(text)
          message = body.message || message
        } catch {
          // Response was not JSON
        }
        throw new Error(message)
      }
      const result: MatchListResult = await res.json()
      matches.value = result.matches ?? []
      matchesTotal.value = result.total
    } catch (e: unknown) {
      matchesError.value = e instanceof Error ? e.message : 'An unexpected error occurred'
    } finally {
      matchesLoading.value = false
    }
  }

  return {
    statsLoading,
    matchesLoading,
    statsError,
    matchesError,
    stats,
    matches,
    matchesTotal,
    searchPlayer,
    fetchStats,
    fetchMatches,
  }
})

export const useSquadStore = defineStore('squad', () => {
  const loading = ref(false)
  const error = ref<string | null>(null)

  return { loading, error }
})

export const useCompareStore = defineStore('compare', () => {
  const players = ref<{ platform: string; gamertag: string }[]>([])

  function addPlayer(platform: string, gamertag: string) {
    if (players.value.length < 4) {
      players.value.push({ platform, gamertag })
    }
  }

  function removePlayer(index: number) {
    players.value.splice(index, 1)
  }

  return { players, addPlayer, removePlayer }
})
