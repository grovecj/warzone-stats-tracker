<script setup lang="ts">
import { NCard, NText, NH2 } from 'naive-ui'
</script>

<template>
  <div class="compare-view">
    <NH2>Compare Players</NH2>
    <NCard>
      <NText :depth="2">Player comparison will be implemented here (issue #14)</NText>
    </NCard>
  </div>
</template>

<style scoped>
.compare-view {
  display: flex;
  flex-direction: column;
  gap: 24px;
}
</style>
