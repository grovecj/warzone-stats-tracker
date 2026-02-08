<script setup lang="ts">
import { computed } from 'vue'
import { VChart, CHART_COLORS, DARK_THEME } from '@/utils/echarts'
import type { PlayerStats } from '@/types'

const props = defineProps<{
  stats: PlayerStats
}>()

const option = computed(() => {
  const s = props.stats
  const headPct = s.kills > 0 ? (s.headshots / s.kills) * 100 : 0
  const survivalPct = s.matchesPlayed > 0 ? (s.topTen / s.matchesPlayed) * 100 : 0
  const aggression = s.matchesPlayed > 0 ? (s.kills / s.matchesPlayed) : 0

  return {
    ...DARK_THEME,
    radar: {
      indicator: [
        { name: 'K/D', max: 4 },
        { name: 'Win%', max: 30 },
        { name: 'Score/Min', max: 200 },
        { name: 'HS%', max: 50 },
        { name: 'Survival', max: 50 },
        { name: 'Aggression', max: 15 },
      ],
      shape: 'polygon' as const,
      splitArea: { areaStyle: { color: ['transparent'] } },
      axisLine: { lineStyle: { color: CHART_COLORS.border } },
      splitLine: { lineStyle: { color: CHART_COLORS.border } },
      axisName: { color: CHART_COLORS.textMuted },
    },
    tooltip: {},
    series: [
      {
        type: 'radar',
        data: [
          {
            value: [
              s.kdRatio,
              s.winPct * 100,
              s.scorePerMin,
              headPct,
              survivalPct,
              aggression,
            ],
            name: 'Stats',
            areaStyle: { color: `${CHART_COLORS.accent}33` },
            lineStyle: { color: CHART_COLORS.accent },
            itemStyle: { color: CHART_COLORS.accent },
          },
        ],
      },
    ],
  }
})
</script>

<template>
  <VChart :option="option" autoresize style="height: 300px" />
</template>
