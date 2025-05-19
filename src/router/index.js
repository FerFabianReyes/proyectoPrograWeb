import { createRouter, createWebHistory } from 'vue-router'
import inicioSecion from '../components/inicioSecion.vue'
import registro from '../components/registro.vue'
import menu from '@/components/menu.vue'
import listaDispo from '@/components/dispositivos/listaDispo.vue'
import listaView from '@/components/dispositivos/viewDispositivos.vue'
import agregarBomba from '@/components/agregarBomba.vue'
import analisisBomba from '@/components/dispositivos/analisisBomba.vue'


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
      children: [
        {
          path: '',
          component: listaView,
          children: [
            {
              path: 'analisisBomba',
              name: 'analisisBomba',
              component: analisisBomba
            }
          ]
        },
        {
          path: 'agregarBomba',
          name: 'agregarBomba',
          component: agregarBomba
        }
      ]
    }
    ,
    {
      path: '/listadispo',
      name: 'listadispo',
      component: listaDispo,
    },
  ],
})

export default router
