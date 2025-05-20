<template>
  <div
    class="card dispositivo"
    @click="irAAnálisis"
    @mouseover="hover = true"
    @mouseleave="hover = false"
    :style="{ boxShadow: hover ? '0 4px 12px rgba(0,0,0,0.3)' : 'none', cursor: 'pointer', transition: 'box-shadow 0.3s ease' }"
  >
    <div class="card-body">
      <h5 class="card-title">{{ dispositivo.nombre }}</h5>
      <p class="card-text">Estado: {{ dispositivo.estado }}</p>
      <p class="card-text">Ubicación: {{ dispositivo.ubicacion }}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'dispositivo',
  props: ['dispositivo'], //recibe un objeto desde el padre
  data() {
    return {
      hover: false //para ver si el mouse está en el componente
    };
  },
  methods: {
    irAAnálisis() {
      this.$router.push({//cambia la ruta y pasa datos
  name: 'analisisBomba',
  query: {
    nombre: this.dispositivo.nombre,
    estado: this.dispositivo.estado,
    ubicacion: this.dispositivo.ubicacion
  }
})

    }
  }
};
</script>

<style scoped>
.card {
  margin: 10px;
  padding: 15px;
  border-radius: 10px;
  border: 1px solid #ddd;
}
</style>
