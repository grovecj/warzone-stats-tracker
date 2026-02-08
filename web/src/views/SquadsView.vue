<script setup lang="ts">
import { NCard, NText, NH2 } from 'naive-ui'
</script>

<template>
  <div class="squads-view">
    <NH2>My Squads</NH2>
    <NCard>
      <NText :depth="2">Squad list will be displayed here (issue #17)</NText>
    </NCard>
  </div>
</template>

<style scoped>
.squads-view {
  display: flex;
  flex-direction: column;
  gap: 24px;
}
</style>
