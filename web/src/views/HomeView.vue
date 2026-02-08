<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { NCard, NInput, NButton, NSelect, NSpace, NH1, NText } from 'naive-ui'

const router = useRouter()
const gamertag = ref('')
const platform = ref('uno')

const platformOptions = [
  { label: 'Activision', value: 'uno' },
  { label: 'Xbox', value: 'xbl' },
  { label: 'PlayStation', value: 'psn' },
  { label: 'Battle.net', value: 'battle' },
  { label: 'Steam', value: 'steam' },
]

function search() {
  if (gamertag.value.trim()) {
    router.push({
      name: 'player',
      params: { platform: platform.value, gamertag: gamertag.value.trim() },
    })
  }
}
</script>

<template>
  <div class="home">
    <div class="hero">
      <NH1 class="hero-title">
        <NText class="hero-accent">WZ</NText> Stats Tracker
      </NH1>
      <NText :depth="2" class="hero-subtitle">
        Search for any player to view their Warzone statistics
      </NText>
    </div>

    <NCard class="search-card">
      <NSpace vertical :size="16">
        <NSelect
          v-model:value="platform"
          :options="platformOptions"
          placeholder="Select platform"
        />
        <NInput
          v-model:value="gamertag"
          placeholder="Enter gamertag..."
          size="large"
          @keyup.enter="search"
        />
        <NButton type="primary" size="large" block @click="search"> Search Player </NButton>
      </NSpace>
    </NCard>
  </div>
</template>

<style scoped>
.home {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 80px;
}

.hero {
  text-align: center;
  margin-bottom: 48px;
}

.hero-title {
  font-size: 48px;
  font-weight: 800;
  margin-bottom: 12px;
}

.hero-accent {
  color: #00e5ff !important;
  font-size: inherit;
  font-weight: inherit;
}

.hero-subtitle {
  font-size: 18px;
}

.search-card {
  width: 100%;
  max-width: 480px;
}
</style>
