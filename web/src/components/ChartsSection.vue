<script setup lang="ts">
import { NCard, NText } from 'naive-ui'
import type { PlayerStats, Match } from '@/types'
import RadarChart from './charts/RadarChart.vue'
import KDGauge from './charts/KDGauge.vue'
import ModeComparisonChart from './charts/ModeComparisonChart.vue'
import PlacementChart from './charts/PlacementChart.vue'
import PerformanceTrend from './charts/PerformanceTrend.vue'

defineProps<{
  stats: PlayerStats
  matches: Match[]
}>()
</script>

<template>
  <div class="charts-section">
    <NCard title="Player Radar">
      <RadarChart :stats="stats" />
    </NCard>

    <NCard title="K/D Gauge">
      <KDGauge :kd-ratio="stats.kdRatio" />
    </NCard>

    <NCard title="Placement Distribution">
      <PlacementChart :stats="stats" />
    </NCard>

    <NCard title="Mode Comparison">
      <template v-if="stats.modeBreakdown && Object.keys(stats.modeBreakdown).length > 0">
        <ModeComparisonChart :mode-breakdown="stats.modeBreakdown" />
      </template>
      <template v-else>
        <NText :depth="3">No mode data available</NText>
      </template>
    </NCard>

    <NCard title="Performance Trend" class="charts-section__full">
      <template v-if="matches.length > 0">
        <PerformanceTrend :matches="matches" />
      </template>
      <template v-else>
        <NText :depth="3">No match data available for trends</NText>
      </template>
    </NCard>
  </div>
</template>

<style scoped>
.charts-section {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.charts-section__full {
  grid-column: 1 / -1;
}

@media (max-width: 768px) {
  .charts-section {
    grid-template-columns: 1fr;
  }
}
</style>
