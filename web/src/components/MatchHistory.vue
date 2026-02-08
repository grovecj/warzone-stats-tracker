<script setup lang="ts">
import { computed, h } from 'vue'
import { NDataTable, NTag, NText } from 'naive-ui'
import type { DataTableColumns } from 'naive-ui'
import type { Match } from '@/types'
import { formatNumber, formatKD } from '@/utils/format'
import { getModeLabel } from '@/utils/modes'

const props = defineProps<{
  matches: Match[]
  loading: boolean
}>()

function formatDateTime(dateStr: string): string {
  const d = new Date(dateStr)
  return d.toLocaleString('en-US', {
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}

const filterOptions = computed(() => {
  const modes = [...new Set(props.matches.map((m) => m.mode))]
  return modes.map((m) => ({ label: getModeLabel(m), value: m }))
})

const columns = computed<DataTableColumns<Match>>(() => [
  {
    title: 'Date/Time',
    key: 'matchTime',
    width: 140,
    render(row) {
      return formatDateTime(row.matchTime)
    },
    sorter: (a, b) => new Date(a.matchTime).getTime() - new Date(b.matchTime).getTime(),
  },
  {
    title: 'Mode',
    key: 'mode',
    width: 140,
    render(row) {
      return getModeLabel(row.mode)
    },
    filter(value, row) {
      return row.mode === value
    },
    filterOptions: filterOptions.value,
  },
  {
    title: 'Map',
    key: 'mapName',
    width: 120,
    render(row) {
      return row.mapName || row.map || '—'
    },
  },
  {
    title: 'Place',
    key: 'placement',
    width: 80,
    sorter: (a, b) => a.placement - b.placement,
    render(row) {
      if (row.placement === 1) {
        return h(NTag, { type: 'warning', size: 'small', bordered: false }, () => '#1')
      }
      return `#${row.placement}`
    },
  },
  {
    title: 'Kills',
    key: 'kills',
    width: 70,
    sorter: (a, b) => a.kills - b.kills,
    render(row) {
      return formatNumber(row.kills)
    },
  },
  {
    title: 'Deaths',
    key: 'deaths',
    width: 80,
    sorter: (a, b) => a.deaths - b.deaths,
    render(row) {
      return formatNumber(row.deaths)
    },
  },
  {
    title: 'K/D',
    key: 'kdRatio',
    width: 70,
    sorter: (a, b) => (a.kdRatio ?? 0) - (b.kdRatio ?? 0),
    render(row) {
      const kd = row.kdRatio ?? (row.deaths > 0 ? row.kills / row.deaths : row.kills)
      return formatKD(kd)
    },
  },
  {
    title: 'Damage',
    key: 'damageDealt',
    width: 90,
    sorter: (a, b) => a.damageDealt - b.damageDealt,
    render(row) {
      return formatNumber(row.damageDealt)
    },
  },
  {
    title: 'Gulag',
    key: 'gulagResult',
    width: 70,
    render(row) {
      if (row.gulagResult === 'win') {
        return h(NText, { style: 'color: #52c41a; font-weight: 600' }, () => 'W')
      }
      if (row.gulagResult === 'loss') {
        return h(NText, { style: 'color: #ff4d4f; font-weight: 600' }, () => 'L')
      }
      return h(NText, { depth: 3 }, () => '—')
    },
  },
])

function rowClassName(row: Match): string {
  if (row.placement === 1) return 'match-win-row'
  return ''
}
</script>

<template>
  <div class="match-history">
    <NDataTable
      :columns="columns"
      :data="matches"
      :loading="loading"
      :row-class-name="rowClassName"
      :bordered="false"
      :single-line="false"
      size="small"
      :pagination="{ pageSize: 20 }"
      :scroll-x="800"
    />
  </div>
</template>

<style scoped>
.match-history {
  width: 100%;
}

:deep(.match-win-row td) {
  background: rgba(250, 173, 20, 0.08) !important;
}
</style>
