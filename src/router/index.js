import { createRouter, createWebHistory } from 'vue-router'
import inicioSecion from '../components/inicioSecion.vue'
import registro from '../components/registro.vue'
import menu from '@/components/menu.vue'
import listaDispo from '@/components/dispositivos/listaDispo.vue'
import listaView from '@/components/dispositivos/viewDispositivos.vue'
import agregarBomba from '@/components/agregarBomba.vue'


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
      children:[
        {
          path: '',
          name: 'listaView',
          component: listaView,
        },
        {
          path: 'agregarBomba',
          name: 'agregarBomba',
          component: agregarBomba
        }
      ]
    },
    {
      path: '/listadispo',
      name: 'listadispo',
      component: listaDispo,
    },
  ],
})

export default router
