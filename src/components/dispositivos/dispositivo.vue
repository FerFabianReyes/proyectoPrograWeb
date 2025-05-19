<template>
  <div class="card mb-3 p-3">
    <h5>{{ dispositivo.nombre }}</h5>
    <p><strong>Ubicación:</strong> {{ dispositivo.ubicacion }}</p>
    <p><strong>Estado:</strong> {{ dispositivo.estado }}</p>
  </div>
</template>

<script>
export default {
  name: 'dispositivo',
  props: {
    dispositivo: {
      type: Object,
      required: true
    }
  }
}
</script>

<style>
.card {
  border: 1px solid #ccc;
  border-radius: 8px;
}
</style>
