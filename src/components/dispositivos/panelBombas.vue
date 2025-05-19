<template>
  <div class="d-flex flex-row">
    <!-- Lista de bombas -->
    <div style="width: 50%; padding: 10px; border-right: 1px solid #ccc;">
      <h2>Lista de Bombas</h2>
      <dispositivo
        v-for="(item, index) in listaDispositivos"
        :key="index"
        :dispositivo="item"
        @click="seleccionarBomba(item)"
      />
    </div>

    <!-- Análisis de bomba seleccionada -->
    <div style="width: 50%; padding: 10px;">
      <analisis-bomba v-if="bombaSeleccionada" :bomba="bombaSeleccionada" />
      <div v-else>
        <p class="text-muted">Selecciona una bomba para ver el análisis</p>
      </div>
    </div>
  </div>
</template>

<script>
import dispositivo from './dispositivo.vue'
import analisisBomba from './analisisBomba.vue'

export default {
  name: 'panelBombas',
  components: { dispositivo, analisisBomba },
  data() {
    return {
      listaDispositivos: [],
      bombaSeleccionada: null
    }
  },
  mounted() {
    const bombas = JSON.parse(localStorage.getItem('bombas')) || [];
    this.listaDispositivos = bombas.map(bomba => ({
      ...bomba
    }));
  },
  methods: {
    seleccionarBomba(bomba) {
      this.bombaSeleccionada = bomba;
    }
  }
}
</script>
