<script setup lang="ts">
import { computed } from 'vue'
import { VChart, CHART_COLORS, DARK_THEME } from '@/utils/echarts'
import type { PlayerStats } from '@/types'

const props = defineProps<{
  stats: PlayerStats
}>()

const option = computed(() => {
  const s = props.stats
  const other = Math.max(
    0,
    s.matchesPlayed - s.wins - (s.topFive - s.wins) - (s.topTen - s.topFive) - (s.topTwentyFive - s.topTen),
  )

  return {
    ...DARK_THEME,
    tooltip: { trigger: 'item' as const },
    legend: {
      orient: 'vertical' as const,
      right: 10,
      top: 'center' as const,
      textStyle: { color: CHART_COLORS.textMuted },
    },
    series: [
      {
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        label: { show: false },
        emphasis: { label: { show: true, fontSize: 14, color: CHART_COLORS.text } },
        data: [
          { value: s.wins, name: 'Wins', itemStyle: { color: CHART_COLORS.cyan } },
          { value: Math.max(0, s.topFive - s.wins), name: 'Top 5', itemStyle: { color: CHART_COLORS.green } },
          { value: Math.max(0, s.topTen - s.topFive), name: 'Top 10', itemStyle: { color: CHART_COLORS.blue } },
          { value: Math.max(0, s.topTwentyFive - s.topTen), name: 'Top 25', itemStyle: { color: CHART_COLORS.purple } },
          { value: other, name: 'Other', itemStyle: { color: CHART_COLORS.border } },
        ],
      },
    ],
  }
})
</script>

<template>
  <VChart :option="option" autoresize style="height: 300px" />
</template>
