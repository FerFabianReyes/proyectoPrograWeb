<template>
  <div class="informacion">
    <div class="list-group">
      <a href="#" class="list-group-item list-group-item-action">
        <div class="d-flex w-100 justify-content-between">
          <h5 class="mb-1">{{ dispositivo.nombre }}</h5>
          <small>{{ dispositivo.estado || 'Sin estado' }}</small>
        </div>
        <small>ID: {{ dispositivo.id }} | Ubicación: {{ dispositivo.ubicacion }}</small>
      </a>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dispositivo',
  props: {
    dispositivo: Object
  }
}
</script>
