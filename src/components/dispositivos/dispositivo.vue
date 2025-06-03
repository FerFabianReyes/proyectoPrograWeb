<template>
  <div
    class="dispositivo"
    @click="seleccionar"
    style="cursor: pointer; border: 1px solid #ccc; padding: 10px; margin: 10px 0; border-radius: 8px; background-color: #f9f9f9; transition: background-color 0.3s;"
    @mouseover="hover = true"
    @mouseleave="hover = false"
    :style="{ backgroundColor: hover ? '#e6f7ff' : '#f9f9f9' }"
  >
    <h3 style="margin: 0 0 10px 0;">{{ dispositivo.nombre }}</h3>
    <p style="margin: 0;">Ubicación: {{ dispositivo.ubicacion }}</p>
    <p style="margin: 0;">Estado: {{ dispositivo.estado }}</p>
  </div>
</template>

<script>
export default {
  props: {
    dispositivo: Object
  },
  data() {
    return {
      hover: false
    };
  },
  methods: {
    seleccionar() {
      this.$emit('seleccionar', this.dispositivo);
    }
  }
}
</script>
