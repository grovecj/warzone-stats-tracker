<script setup lang="ts">
import { computed } from 'vue'
import { VChart, CHART_COLORS, DARK_THEME } from '@/utils/echarts'
import type { Match } from '@/types'

const props = defineProps<{
  matches: Match[]
}>()

const option = computed(() => {
  // Group matches by date and compute average K/D
  const byDate = new Map<string, { totalKD: number; count: number }>()
  for (const m of props.matches) {
    const date = new Date(m.matchTime).toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
    const kd = m.kdRatio ?? (m.deaths > 0 ? m.kills / m.deaths : m.kills)
    const entry = byDate.get(date) || { totalKD: 0, count: 0 }
    entry.totalKD += kd
    entry.count++
    byDate.set(date, entry)
  }

  // Sort chronologically (matches are newest-first, so reverse)
  const entries = [...byDate.entries()].reverse()
  const dates = entries.map(([d]) => d)
  const avgKDs = entries.map(([, v]) => Number((v.totalKD / v.count).toFixed(2)))

  return {
    ...DARK_THEME,
    tooltip: { trigger: 'axis' as const },
    grid: { left: '3%', right: '4%', bottom: '3%', containLabel: true },
    xAxis: {
      type: 'category' as const,
      data: dates,
      axisLabel: { color: CHART_COLORS.textMuted },
      axisLine: { lineStyle: { color: CHART_COLORS.border } },
    },
    yAxis: {
      type: 'value' as const,
      name: 'Avg K/D',
      axisLabel: { color: CHART_COLORS.textMuted },
      splitLine: { lineStyle: { color: CHART_COLORS.border } },
      nameTextStyle: { color: CHART_COLORS.textMuted },
    },
    series: [
      {
        name: 'Avg K/D',
        type: 'line',
        data: avgKDs,
        smooth: true,
        lineStyle: { color: CHART_COLORS.accent },
        itemStyle: { color: CHART_COLORS.accent },
        areaStyle: { color: `${CHART_COLORS.accent}22` },
      },
    ],
  }
})
</script>

<template>
  <VChart v-if="matches.length > 0" :option="option" autoresize style="height: 300px" />
</template>
