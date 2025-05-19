<template>
  <div>
    <h1>Lista de dispositivos</h1> <br>
    <dispositivo
      v-for="(item, index) in listaDispositivos"
      :key="index"
      :dispositivo="item"
    />
  </div>
</template>

<script>
import dispositivo from './dispositivo.vue'

export default {
  name: 'viewDispositivos',
  components: { dispositivo },
  data() {
    return {
      listaDispositivos: []
    }
  },
  mounted() {
    const bombas = JSON.parse(localStorage.getItem('bombas')) || [];
    this.listaDispositivos = bombas.map(bomba => ({
      nombre: bomba.nombre,
      estado: bomba.estado,
      ubicacion: bomba.coordenadas // usamos "ubicacion" como alias de "coordenadas"
    }));
  }
}
</script>
