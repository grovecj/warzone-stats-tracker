export interface ModeInfo {
  key: string
  label: string
  icon: string
}

export const MODES: ModeInfo[] = [
  { key: 'wz', label: 'All Modes', icon: 'GI' },
  { key: 'br', label: 'Battle Royale', icon: 'BR' },
  { key: 'rebirth', label: 'Resurgence', icon: 'RS' },
  { key: 'ranked', label: 'Ranked', icon: 'RK' },
  { key: 'plunder', label: 'Plunder', icon: 'PL' },
]

export const MODE_LABELS: Record<string, string> = {
  br_all: 'Battle Royale',
  br_brsolo: 'BR Solos',
  br_brduos: 'BR Duos',
  br_brtrios: 'BR Trios',
  br_brquads: 'BR Quads',
  br_rebirth_rbrthduos: 'Resurgence Duos',
  br_rebirth_rbrthtrios: 'Resurgence Trios',
  br_rebirth_rbrthquads: 'Resurgence Quads',
  br_dmz: 'Plunder',
  br_plnbld: 'Blood Money',
  br_kingslayer_kingsltrios: 'King Slayer',
  br_mini_miniroyale: 'Mini Royale',
}

export function getModeLabel(mode: string): string {
  return MODE_LABELS[mode] || mode.replace(/^br_/, '').replace(/_/g, ' ')
}
