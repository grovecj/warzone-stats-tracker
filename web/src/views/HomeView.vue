<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { NCard, NInput, NButton, NSelect, NSpace, NH1, NText, NAlert } from 'naive-ui'
import { usePlayerStore } from '@/stores'

const router = useRouter()
const playerStore = usePlayerStore()

const gamertag = ref('')
const platform = ref('uno')
const game = ref('wz')
const searching = ref(false)
const searchError = ref<string | null>(null)

interface RecentSearch {
  platform: string
  gamertag: string
  game?: string
  timestamp: number
}

const recentSearches = ref<RecentSearch[]>([])

const platformOptions = [
  { label: 'Activision', value: 'uno' },
  { label: 'Xbox', value: 'xbl' },
  { label: 'PlayStation', value: 'psn' },
  { label: 'Battle.net', value: 'battle' },
  { label: 'Steam', value: 'steam' },
]

const gameOptions = [
  { label: 'Warzone', value: 'wz' },
  { label: 'Warzone 2', value: 'wz2' },
  { label: 'MW Multiplayer', value: 'mw-mp' },
  { label: 'MW2 Multiplayer', value: 'mw2-mp' },
]

const gameLabels: Record<string, string> = {
  wz: 'WZ',
  wz2: 'WZ2',
  'mw-mp': 'MW MP',
  'mw2-mp': 'MW2 MP',
}

function gameToTitleMode(g: string): { title: string; mode: string } {
  switch (g) {
    case 'wz2':
      return { title: 'mw2', mode: 'wz2' }
    case 'mw-mp':
      return { title: 'mw', mode: 'mp' }
    case 'mw2-mp':
      return { title: 'mw2', mode: 'mp' }
    default:
      return { title: 'mw', mode: 'wz' }
  }
}

const platformLabels: Record<string, string> = {
  uno: 'Activision',
  xbl: 'Xbox',
  psn: 'PlayStation',
  battle: 'Battle.net',
  steam: 'Steam',
}

const demoPlayers = [
  { gamertag: 'TacticalNuke99', platform: 'xbl', label: 'Elite' },
  { gamertag: 'ShadowSniper_TTV', platform: 'psn', label: 'Skilled' },
  { gamertag: 'GhostRecon42', platform: 'uno', label: 'Above Avg' },
  { gamertag: 'CasualCarl', platform: 'xbl', label: 'Average' },
  { gamertag: 'NoobMaster69', platform: 'battle', label: 'Casual' },
]

function goToDemo(demo: { gamertag: string; platform: string }) {
  platform.value = demo.platform
  gamertag.value = demo.gamertag
  game.value = 'wz'
  search()
}

onMounted(() => {
  loadRecentSearches()
})

function loadRecentSearches() {
  try {
    const stored = localStorage.getItem('wz-recent-searches')
    if (stored) {
      recentSearches.value = JSON.parse(stored)
    }
  } catch {
    // Ignore parse errors
  }
}

function saveRecentSearch(plat: string, tag: string, g: string) {
  const filtered = recentSearches.value.filter(
    (s) => !(s.platform === plat && s.gamertag === tag && s.game === g),
  )
  filtered.unshift({ platform: plat, gamertag: tag, game: g, timestamp: Date.now() })
  recentSearches.value = filtered.slice(0, 5)
  localStorage.setItem('wz-recent-searches', JSON.stringify(recentSearches.value))
}

async function search() {
  const tag = gamertag.value.trim()
  if (!tag) return

  searching.value = true
  searchError.value = null

  const { title, mode } = gameToTitleMode(game.value)

  try {
    await playerStore.searchPlayer(platform.value, tag, title, mode)
    saveRecentSearch(platform.value, tag, game.value)
    router.push({
      name: 'player',
      params: { platform: platform.value, gamertag: tag },
      query: { title, mode },
    })
  } catch (e: unknown) {
    searchError.value = e instanceof Error ? e.message : 'An unexpected error occurred'
  } finally {
    searching.value = false
  }
}

function goToRecent(recent: RecentSearch) {
  platform.value = recent.platform
  gamertag.value = recent.gamertag
  game.value = recent.game || 'wz'
  search()
}
</script>

<template>
  <div class="home">
    <div class="hero">
      <NH1 class="hero-title">
        <NText class="hero-accent">WZ</NText> Stats Tracker
      </NH1>
      <NText :depth="2" class="hero-subtitle">
        Search for any player to view their Warzone statistics
      </NText>
    </div>

    <NCard class="search-card">
      <NSpace vertical :size="16">
        <NAlert v-if="searchError" type="error" closable @close="searchError = null">
          {{ searchError }}
        </NAlert>
        <NSelect
          v-model:value="game"
          :options="gameOptions"
          placeholder="Select game"
        />
        <NSelect
          v-model:value="platform"
          :options="platformOptions"
          placeholder="Select platform"
        />
        <NInput
          v-model:value="gamertag"
          placeholder="Enter gamertag..."
          size="large"
          :disabled="searching"
          @keyup.enter="search"
        />
        <NButton
          type="primary"
          size="large"
          block
          :loading="searching"
          :disabled="!gamertag.trim()"
          @click="search"
        >
          Search Player
        </NButton>
      </NSpace>
    </NCard>

    <div class="demo-players">
      <NText :depth="3" class="section-title">Demo Players</NText>
      <div class="demo-list">
        <NButton
          v-for="demo in demoPlayers"
          :key="demo.gamertag"
          quaternary
          size="small"
          @click="goToDemo(demo)"
        >
          {{ demo.gamertag }}
          <NText :depth="3" style="margin-left: 4px; font-size: 11px">
            {{ platformLabels[demo.platform] }} &middot; {{ demo.label }}
          </NText>
        </NButton>
      </div>
    </div>

    <div v-if="recentSearches.length > 0" class="recent-searches">
      <NText :depth="3" class="section-title">Recent Searches</NText>
      <div class="recent-list">
        <NButton
          v-for="recent in recentSearches"
          :key="`${recent.platform}-${recent.gamertag}`"
          quaternary
          size="small"
          @click="goToRecent(recent)"
        >
          {{ recent.gamertag }}
          <NText :depth="3" style="margin-left: 4px; font-size: 11px">
            {{ platformLabels[recent.platform] || recent.platform }}
            {{ recent.game ? `/ ${gameLabels[recent.game] || recent.game}` : '' }}
          </NText>
        </NButton>
      </div>
    </div>
  </div>
</template>

<style scoped>
.home {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 80px;
}

.hero {
  text-align: center;
  margin-bottom: 48px;
}

.hero-title {
  font-size: 48px;
  font-weight: 800;
  margin-bottom: 12px;
}

.hero-accent {
  color: #00e5ff !important;
  font-size: inherit;
  font-weight: inherit;
}

.hero-subtitle {
  font-size: 18px;
}

.search-card {
  width: 100%;
  max-width: 480px;
}

.demo-players,
.recent-searches {
  margin-top: 32px;
  width: 100%;
  max-width: 480px;
  text-align: center;
}

.section-title {
  font-size: 12px;
  text-transform: uppercase;
  letter-spacing: 1px;
  display: block;
  margin-bottom: 12px;
}

.demo-list,
.recent-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  justify-content: center;
}
</style>
