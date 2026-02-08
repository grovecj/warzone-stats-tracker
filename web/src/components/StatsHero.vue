<script setup lang="ts">
import type { PlayerStats } from '@/types'
import { formatKD, formatNumber, kdColorClass } from '@/utils/format'
import StatCard from './StatCard.vue'

defineProps<{
  stats: PlayerStats
}>()
</script>

<template>
  <div class="stats-hero">
    <StatCard label="K/D Ratio" :value="formatKD(stats.kdRatio)" :color-class="kdColorClass(stats.kdRatio)" size="lg" />
    <StatCard label="Wins" :value="formatNumber(stats.wins)" size="lg" />
    <StatCard label="Total Kills" :value="formatNumber(stats.kills)" size="lg" />
  </div>
</template>

<style scoped>
.stats-hero {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

@media (max-width: 640px) {
  .stats-hero {
    grid-template-columns: 1fr;
  }
}
</style>
