import { createRouter, createWebHistory } from 'vue-router'
import inicioSecion from '../components/inicioSecion.vue'
import registro from '../components/registro.vue'
import menu from '@/components/menu.vue'
import listaDispo from '@/components/dispositivos/listaDispo.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'inicio',
      component: inicioSecion,
    },
    {
      path: '/registro',
      name: 'registro',
      component: registro,
    },
    {
      path: '/menu',
      name: 'menu',
      component: menu,
    },
    {
      path: '/listadispo',
      name: 'listadispo',
      component: listaDispo,
    },
  ],
})

export default router
