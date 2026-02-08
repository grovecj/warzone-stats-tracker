<script setup lang="ts">
import { useRoute } from 'vue-router'
import { NCard, NText, NH2 } from 'naive-ui'

const route = useRoute()
const squadId = route.params.id as string
</script>

<template>
  <div class="squad-view">
    <NH2>Squad Dashboard</NH2>
    <NCard>
      <NText :depth="2">Squad {{ squadId }} dashboard will be implemented here (issue #17)</NText>
    </NCard>
  </div>
</template>

<style scoped>
.squad-view {
  display: flex;
  flex-direction: column;
  gap: 24px;
}
</style>
