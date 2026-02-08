<script setup lang="ts">
import { NCard, NText } from 'naive-ui'
import type { ModeStats } from '@/types'
import { formatKD, formatNumber, kdColorClass } from '@/utils/format'
import { getModeLabel } from '@/utils/modes'

defineProps<{
  modeBreakdown: Record<string, ModeStats>
}>()

const emit = defineEmits<{
  selectMode: [mode: string]
}>()
</script>

<template>
  <div class="mode-summary-row">
    <NCard
      v-for="(modeStats, mode) in modeBreakdown"
      :key="mode"
      class="mode-card"
      hoverable
      @click="emit('selectMode', String(mode))"
    >
      <div class="mode-card__content">
        <NText :depth="2" class="mode-card__name">{{ getModeLabel(String(mode)) }}</NText>
        <NText :class="['mode-card__kd', kdColorClass(modeStats.kdRatio)]">
          {{ formatKD(modeStats.kdRatio) }} K/D
        </NText>
        <NText :depth="3" class="mode-card__detail">
          {{ formatNumber(modeStats.wins) }} W &middot; {{ formatNumber(modeStats.matchesPlayed) }} G
        </NText>
      </div>
    </NCard>
  </div>
</template>

<style scoped>
.mode-summary-row {
  display: flex;
  gap: 12px;
  overflow-x: auto;
  padding-bottom: 4px;
}

.mode-card {
  min-width: 160px;
  flex-shrink: 0;
  cursor: pointer;
}

.mode-card__content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.mode-card__name {
  font-size: 12px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.mode-card__kd {
  font-size: 20px;
  font-weight: 700;
}

.mode-card__detail {
  font-size: 12px;
}

:deep(.kd-excellent) {
  color: #00e5ff !important;
}
:deep(.kd-good) {
  color: #52c41a !important;
}
:deep(.kd-average) {
  color: #faad14 !important;
}
:deep(.kd-poor) {
  color: #ff4d4f !important;
}
</style>
