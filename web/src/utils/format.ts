const numberFormatter = new Intl.NumberFormat('en-US')

export function formatNumber(n: number): string {
  return numberFormatter.format(n)
}

export function formatKD(kd: number): string {
  return kd.toFixed(2)
}

export function formatTimePlayed(seconds: number): string {
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  return `${hours}h ${minutes}m`
}

export function formatPercent(pct: number): string {
  return `${(pct * 100).toFixed(1)}%`
}

export function kdColorClass(kd: number): string {
  if (kd >= 2.0) return 'kd-excellent'
  if (kd >= 1.2) return 'kd-good'
  if (kd >= 0.8) return 'kd-average'
  return 'kd-poor'
}
