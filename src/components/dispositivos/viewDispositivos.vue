<template>
  <div class="lista-view-container" style="display: flex; gap: 20px; height: 80vh;">
    <!-- Contenedor con scroll vertical -->
    <div class="lista-dispositivos"
      style="flex: 1; min-width: 400px; border-right: 1px solid #ccc; padding-right: 20px; overflow-y: auto;">
      <h1>Lista de dispositivos</h1>
      <dispositivo
        v-for="(item, index) in listaDispositivos"
        :key="index"
        :dispositivo="item"
        @seleccionar="seleccionarBomba"
      />
    </div>

    <div class="analisis-bomba" style="flex: 1; min-width: 800px;">
      <analisis-bomba
        v-if="bombaSeleccionada"
        :bomba="bombaSeleccionada"
      />
    </div>
  </div>
</template>

<script>
import dispositivo from './dispositivo.vue'
import analisisBomba from './analisisBomba.vue'

export default {
  name: 'listaView',
  components: { dispositivo, analisisBomba },
  data() {
    return {
      listaDispositivos: [],
      bombaSeleccionada: null
    }
  },
  mounted() {
    const usuarioActivo = JSON.parse(localStorage.getItem('usuarioActivo'));
    const bombas = usuarioActivo?.bombas || [];
    this.listaDispositivos = bombas;
  },
  methods: {
    seleccionarBomba(bomba) {
      this.bombaSeleccionada = bomba;
    }
  }
}
</script>
