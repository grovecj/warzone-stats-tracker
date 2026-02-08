<script setup lang="ts">
import { computed } from 'vue'
import { VChart, CHART_COLORS, DARK_THEME } from '@/utils/echarts'

const props = defineProps<{
  kdRatio: number
}>()

const option = computed(() => ({
  ...DARK_THEME,
  series: [
    {
      type: 'gauge',
      min: 0,
      max: 4,
      splitNumber: 4,
      axisLine: {
        lineStyle: {
          width: 12,
          color: [
            [0.2, CHART_COLORS.red],
            [0.3, CHART_COLORS.yellow],
            [0.5, CHART_COLORS.green],
            [1, CHART_COLORS.cyan],
          ] as [number, string][],
        },
      },
      pointer: { itemStyle: { color: 'auto' }, width: 4 },
      axisTick: { show: false },
      splitLine: { length: 8, lineStyle: { color: CHART_COLORS.border } },
      axisLabel: { color: CHART_COLORS.textMuted, distance: 18, fontSize: 11 },
      title: { show: false },
      detail: {
        fontSize: 28,
        fontWeight: 'bold' as const,
        color: 'auto',
        offsetCenter: [0, '70%'],
        formatter: '{value}',
      },
      data: [{ value: Number(props.kdRatio.toFixed(2)) }],
    },
  ],
}))
</script>

<template>
  <VChart :option="option" autoresize style="height: 250px" />
</template>
