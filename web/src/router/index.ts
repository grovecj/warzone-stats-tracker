import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/player/:platform/:gamertag',
      name: 'player',
      component: () => import('../views/PlayerView.vue'),
    },
    {
      path: '/compare',
      name: 'compare',
      component: () => import('../views/CompareView.vue'),
    },
    {
      path: '/squads',
      name: 'squads',
      component: () => import('../views/SquadsView.vue'),
    },
    {
      path: '/squad/:id',
      name: 'squad',
      component: () => import('../views/SquadView.vue'),
    },
  ],
})

export default router
