<template>
  <div class="container text-center">
    <div class="row">
      <div class="col">
        <p class="h4">Nombre: {{ bomba.nombre }}</p>
      </div>
      <div class="col">
        <p class="h4">Estado: {{ bomba.estado }}</p>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <p class="h5">Ubicación: {{ bomba.ubicacion }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'analisisBomba',
  data() {
    return {
      bomba: {
        nombre: '',
        estado: '',
        ubicacion: ''
      }
    }
  },
  mounted() {
    const query = this.$route.query;
    this.bomba = {
      nombre: query.nombre || '',
      estado: query.estado || '',
      ubicacion: query.ubicacion || ''
    };
  }
}
</script>
