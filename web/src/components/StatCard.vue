<script setup lang="ts">
import { NCard, NText } from 'naive-ui'

defineProps<{
  label: string
  value: string
  colorClass?: string
  size?: 'sm' | 'lg'
}>()
</script>

<template>
  <NCard class="stat-card" :class="[size === 'lg' ? 'stat-card--lg' : 'stat-card--sm']">
    <div class="stat-card__content">
      <NText :depth="3" class="stat-card__label">{{ label }}</NText>
      <NText :class="['stat-card__value', colorClass]">{{ value }}</NText>
    </div>
  </NCard>
</template>

<style scoped>
.stat-card {
  text-align: center;
}

.stat-card__content {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.stat-card__label {
  font-size: 12px;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.stat-card--lg .stat-card__value {
  font-size: 32px;
  font-weight: 700;
}

.stat-card--sm .stat-card__value {
  font-size: 20px;
  font-weight: 600;
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
