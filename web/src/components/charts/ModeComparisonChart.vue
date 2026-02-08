<script setup lang="ts">
import { computed } from 'vue'
import { VChart, CHART_COLORS, DARK_THEME } from '@/utils/echarts'
import type { ModeStats } from '@/types'
import { getModeLabel } from '@/utils/modes'

const props = defineProps<{
  modeBreakdown: Record<string, ModeStats>
}>()

const option = computed(() => {
  const modes = Object.keys(props.modeBreakdown)
  const labels = modes.map(getModeLabel)
  const kills = modes.map((m) => props.modeBreakdown[m]?.kills ?? 0)
  const deaths = modes.map((m) => props.modeBreakdown[m]?.deaths ?? 0)

  return {
    ...DARK_THEME,
    tooltip: { trigger: 'axis' as const, axisPointer: { type: 'shadow' as const } },
    legend: { data: ['Kills', 'Deaths'], textStyle: { color: CHART_COLORS.textMuted } },
    grid: { left: '3%', right: '4%', bottom: '3%', containLabel: true },
    xAxis: {
      type: 'category' as const,
      data: labels,
      axisLabel: { color: CHART_COLORS.textMuted, rotate: labels.length > 5 ? 30 : 0 },
      axisLine: { lineStyle: { color: CHART_COLORS.border } },
    },
    yAxis: {
      type: 'value' as const,
      axisLabel: { color: CHART_COLORS.textMuted },
      splitLine: { lineStyle: { color: CHART_COLORS.border } },
    },
    series: [
      {
        name: 'Kills',
        type: 'bar',
        data: kills,
        itemStyle: { color: CHART_COLORS.accent },
      },
      {
        name: 'Deaths',
        type: 'bar',
        data: deaths,
        itemStyle: { color: CHART_COLORS.red },
      },
    ],
  }
})
</script>

<template>
  <VChart :option="option" autoresize style="height: 300px" />
</template>
