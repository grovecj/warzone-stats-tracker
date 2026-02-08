<script setup lang="ts">
import { useRoute } from 'vue-router'
import { NCard, NText, NH2, NSpin } from 'naive-ui'

const route = useRoute()
const platform = route.params.platform as string
const gamertag = route.params.gamertag as string
</script>

<template>
  <div class="player-view">
    <NH2>
      <NText>{{ gamertag }}</NText>
      <NText :depth="3"> ({{ platform }})</NText>
    </NH2>
    <NCard>
      <NSpin :show="false">
        <NText :depth="2">Player stats will be displayed here (issue #10)</NText>
      </NSpin>
    </NCard>
  </div>
</template>

<style scoped>
.player-view {
  display: flex;
  flex-direction: column;
  gap: 24px;
}
</style>
