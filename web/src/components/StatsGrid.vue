<script setup lang="ts">
import type { PlayerStats } from '@/types'
import { formatNumber, formatKD, formatPercent, formatTimePlayed } from '@/utils/format'
import StatCard from './StatCard.vue'

defineProps<{
  stats: PlayerStats
}>()
</script>

<template>
  <div class="stats-grid">
    <StatCard label="Deaths" :value="formatNumber(stats.deaths)" size="sm" />
    <StatCard label="Win %" :value="formatPercent(stats.winPct)" size="sm" />
    <StatCard label="Score/Min" :value="formatKD(stats.scorePerMin)" size="sm" />
    <StatCard label="Headshots" :value="formatNumber(stats.headshots)" size="sm" />
    <StatCard label="Time Played" :value="formatTimePlayed(stats.timePlayed)" size="sm" />
    <StatCard label="Matches" :value="formatNumber(stats.matchesPlayed)" size="sm" />
    <StatCard label="Top 5" :value="formatNumber(stats.topFive)" size="sm" />
    <StatCard label="Top 10" :value="formatNumber(stats.topTen)" size="sm" />
    <StatCard label="Top 25" :value="formatNumber(stats.topTwentyFive)" size="sm" />
    <StatCard label="Assists" :value="formatNumber(stats.assists)" size="sm" />
    <StatCard label="Damage Done" :value="formatNumber(stats.damageDone)" size="sm" />
    <StatCard
      label="Level"
      :value="stats.prestige > 0 ? `P${stats.prestige} Lv${stats.level}` : `Lv ${stats.level}`"
      size="sm"
    />
  </div>
</template>

<style scoped>
.stats-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 12px;
}

@media (max-width: 900px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 480px) {
  .stats-grid {
    grid-template-columns: 1fr;
  }
}
</style>
