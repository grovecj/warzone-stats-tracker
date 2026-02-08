<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { NText, NH2, NAlert, NSkeleton, NTabs, NTabPane } from 'naive-ui'
import { usePlayerStore } from '@/stores'
import { MODES } from '@/utils/modes'
import StatsHero from '@/components/StatsHero.vue'
import StatsGrid from '@/components/StatsGrid.vue'
import ChartsSection from '@/components/ChartsSection.vue'
import MatchHistory from '@/components/MatchHistory.vue'
import ModeSummaryRow from '@/components/ModeSummaryRow.vue'

const route = useRoute()
const playerStore = usePlayerStore()

const platform = computed(() => route.params.platform as string)
const gamertag = computed(() => route.params.gamertag as string)
const title = computed(() => (route.query.title as string) || 'mw')
const mode = computed(() => (route.query.mode as string) || 'wz')
const selectedMode = ref(mode.value)

function loadData() {
  playerStore.fetchStats(platform.value, gamertag.value, title.value, selectedMode.value)
  playerStore.fetchMatches(platform.value, gamertag.value, title.value, mode.value)
}

onMounted(loadData)

watch(
  () => [route.params.platform, route.params.gamertag, route.query.title, route.query.mode],
  () => {
    selectedMode.value = mode.value
    loadData()
  },
)

function onModeChange(m: string) {
  selectedMode.value = m
  playerStore.fetchStats(platform.value, gamertag.value, title.value, m)
}

function onModeSummarySelect(m: string) {
  selectedMode.value = m
  playerStore.fetchStats(platform.value, gamertag.value, title.value, m)
}

const modeTabPanes = MODES
</script>

<template>
  <div class="player-view">
    <!-- Player header -->
    <div class="player-header">
      <NH2 style="margin: 0">
        <NText>{{ gamertag }}</NText>
        <NText :depth="3" style="font-size: 16px; margin-left: 8px">({{ platform }})</NText>
      </NH2>
    </div>

    <!-- Error alerts -->
    <NAlert v-if="playerStore.statsError" type="error" closable>
      {{ playerStore.statsError }}
    </NAlert>
    <NAlert v-if="playerStore.matchesError" type="warning" closable>
      {{ playerStore.matchesError }}
    </NAlert>

    <!-- Mode tabs -->
    <NTabs
      :value="selectedMode"
      type="segment"
      animated
      @update:value="onModeChange"
    >
      <NTabPane
        v-for="mode in modeTabPanes"
        :key="mode.key"
        :name="mode.key"
        :tab="mode.label"
      />
    </NTabs>

    <!-- Stats loading skeleton -->
    <template v-if="playerStore.statsLoading">
      <div class="skeleton-hero">
        <NSkeleton height="120px" />
        <NSkeleton height="120px" />
        <NSkeleton height="120px" />
      </div>
      <div class="skeleton-grid">
        <NSkeleton v-for="i in 8" :key="i" height="80px" />
      </div>
    </template>

    <!-- Stats display -->
    <template v-else-if="playerStore.stats">
      <!-- Mode summary row -->
      <ModeSummaryRow
        v-if="playerStore.stats.modeBreakdown && Object.keys(playerStore.stats.modeBreakdown).length > 0"
        :mode-breakdown="playerStore.stats.modeBreakdown"
        @select-mode="onModeSummarySelect"
      />

      <StatsHero :stats="playerStore.stats" />
      <StatsGrid :stats="playerStore.stats" />

      <!-- Charts -->
      <ChartsSection :stats="playerStore.stats" :matches="playerStore.matches" />
    </template>

    <!-- Match history -->
    <div class="section-header">
      <NText :depth="1" style="font-size: 18px; font-weight: 600">Match History</NText>
    </div>
    <MatchHistory :matches="playerStore.matches" :loading="playerStore.matchesLoading" />
  </div>
</template>

<style scoped>
.player-view {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.player-header {
  display: flex;
  align-items: center;
  gap: 12px;
}

.section-header {
  margin-top: 8px;
}

.skeleton-hero {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.skeleton-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 12px;
}

@media (max-width: 640px) {
  .skeleton-hero {
    grid-template-columns: 1fr;
  }
  .skeleton-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>
