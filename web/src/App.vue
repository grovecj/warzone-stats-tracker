<script setup lang="ts">
import { ref } from 'vue'
import { darkTheme, type GlobalThemeOverrides } from 'naive-ui'
import {
  NConfigProvider,
  NLayout,
  NLayoutHeader,
  NLayoutContent,
  NSpace,
  NText,
  NAlert,
} from 'naive-ui'
import { RouterLink, RouterView } from 'vue-router'

const showBanner = ref(sessionStorage.getItem('wz-hide-banner') !== '1')

function dismissBanner() {
  showBanner.value = false
  sessionStorage.setItem('wz-hide-banner', '1')
}

const themeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: '#00e5ff',
    primaryColorHover: '#18ffff',
    primaryColorPressed: '#00b8d4',
    primaryColorSuppl: '#00e5ff',
    bodyColor: '#0a0a0f',
    cardColor: '#12121a',
    modalColor: '#12121a',
    popoverColor: '#12121a',
    tableColor: '#12121a',
    inputColor: '#1a1a28',
    borderColor: '#2a2a3a',
    textColorBase: '#e0e0e0',
    textColor1: '#ffffff',
    textColor2: '#b0b0c0',
    textColor3: '#707088',
  },
}
</script>

<template>
  <NConfigProvider :theme="darkTheme" :theme-overrides="themeOverrides">
    <NLayout class="app-layout">
      <NLayoutHeader class="app-header" bordered>
        <div class="header-content">
          <RouterLink to="/" class="logo">
            <NText :depth="1" class="logo-text">WZ Stats</NText>
          </RouterLink>
          <NSpace :size="24" align="center">
            <RouterLink to="/" class="nav-link">
              <NText :depth="2">Home</NText>
            </RouterLink>
            <RouterLink to="/compare" class="nav-link">
              <NText :depth="2">Compare</NText>
            </RouterLink>
            <RouterLink to="/squads" class="nav-link">
              <NText :depth="2">Squads</NText>
            </RouterLink>
          </NSpace>
        </div>
      </NLayoutHeader>
      <div v-if="showBanner" class="demo-banner-wrapper">
        <NAlert type="warning" closable @close="dismissBanner">
          <strong>Demo Mode</strong> — The Activision stats API was shut down in 2024.
          All player data shown here is generated sample data for demonstration purposes.
        </NAlert>
      </div>
      <NLayoutContent class="app-content" content-style="padding: 24px;">
        <RouterView />
      </NLayoutContent>
    </NLayout>
  </NConfigProvider>
</template>

<style scoped>
.app-layout {
  min-height: 100vh;
}

.app-header {
  padding: 0 24px;
  background: #0d0d14;
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 56px;
  max-width: 1200px;
  margin: 0 auto;
}

.logo {
  text-decoration: none;
}

.logo-text {
  font-size: 20px;
  font-weight: 700;
  letter-spacing: 2px;
  color: #00e5ff !important;
}

.nav-link {
  text-decoration: none;
  transition: opacity 0.2s;
}

.nav-link:hover {
  opacity: 0.8;
}

.demo-banner-wrapper {
  max-width: 1200px;
  margin: 0 auto;
  padding: 12px 24px 0;
  width: 100%;
  box-sizing: border-box;
}

.app-content {
  max-width: 1200px;
  margin: 0 auto;
  width: 100%;
}
</style>
