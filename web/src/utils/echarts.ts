import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { RadarChart, GaugeChart, BarChart, PieChart, LineChart } from 'echarts/charts'
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent,
  RadarComponent,
} from 'echarts/components'
import VChart from 'vue-echarts'

use([
  CanvasRenderer,
  RadarChart,
  GaugeChart,
  BarChart,
  PieChart,
  LineChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent,
  RadarComponent,
])

export const CHART_COLORS = {
  accent: '#00e5ff',
  accentHover: '#18ffff',
  bg: '#0a0a0f',
  cardBg: '#12121a',
  text: '#e0e0e0',
  textMuted: '#707088',
  border: '#2a2a3a',
  red: '#ff4d4f',
  yellow: '#faad14',
  green: '#52c41a',
  cyan: '#00e5ff',
  blue: '#1890ff',
  purple: '#722ed1',
  orange: '#fa8c16',
  series: ['#00e5ff', '#ff4d4f', '#52c41a', '#faad14', '#1890ff', '#722ed1'],
}

export const DARK_THEME = {
  backgroundColor: 'transparent',
  textStyle: { color: CHART_COLORS.text },
  title: { textStyle: { color: CHART_COLORS.text } },
  legend: { textStyle: { color: CHART_COLORS.textMuted } },
}

export { VChart }
