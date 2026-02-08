import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { PlayerStats, Match } from '@/types'

export const usePlayerStore = defineStore('player', () => {
  const loading = ref(false)
  const error = ref<string | null>(null)
  const stats = ref<PlayerStats | null>(null)
  const matches = ref<Match[]>([])

  async function fetchStats(platform: string, gamertag: string, mode = 'wz') {
    loading.value = true
    error.value = null
    try {
      const res = await fetch(`/api/v1/players/${platform}/${gamertag}/stats?mode=${mode}`)
      if (!res.ok) {
        const body = await res.json()
        throw new Error(body.message || 'Failed to fetch stats')
      }
      stats.value = await res.json()
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : 'An unexpected error occurred'
    } finally {
      loading.value = false
    }
  }

  async function fetchMatches(platform: string, gamertag: string) {
    loading.value = true
    error.value = null
    try {
      const res = await fetch(`/api/v1/players/${platform}/${gamertag}/matches`)
      if (!res.ok) {
        const body = await res.json()
        throw new Error(body.message || 'Failed to fetch matches')
      }
      matches.value = await res.json()
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : 'An unexpected error occurred'
    } finally {
      loading.value = false
    }
  }

  return { loading, error, stats, matches, fetchStats, fetchMatches }
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
