<script setup lang="ts">
import { NTag, NText } from 'naive-ui'

defineProps<{
  division?: string
  sr?: number
}>()
</script>

<template>
  <div v-if="division || sr" class="ranked-badge">
    <NTag type="info" size="medium" round>
      <NText style="font-weight: 600">
        {{ division || 'Ranked' }}
        <template v-if="sr"> &middot; {{ sr }} SR</template>
      </NText>
    </NTag>
  </div>
</template>

<style scoped>
.ranked-badge {
  display: inline-flex;
}
</style>
