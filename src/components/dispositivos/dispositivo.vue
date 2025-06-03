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

    <!-- Indicadores de estado en tiempo real -->
    <div style="margin-top: 10px; font-size: 0.9em;">
      <div style="display: flex; justify-content: space-between; margin-bottom: 5px;">
        <span>Potencia:</span>
        <span :class="getStatusClass('potencia')" style="font-weight: bold;">
          {{ dispositivo.valor?.potencia || '---' }} {{ dispositivo.potencia?.unidades || 'KW' }}
        </span>
      </div>
      <div style="display: flex; justify-content: space-between; margin-bottom: 5px;">
        <span>Voltaje:</span>
        <span :class="getStatusClass('voltaje')" style="font-weight: bold;">
          {{ dispositivo.valor?.voltaje || '---' }} {{ dispositivo.voltaje?.unidades || 'V' }}
        </span>
      </div>
      <div style="display: flex; justify-content: space-between; margin-bottom: 5px;">
        <span>Corriente:</span>
        <span :class="getStatusClass('corriente')" style="font-weight: bold;">
          {{ dispositivo.valor?.corriente || '---' }} {{ dispositivo.corriente?.unidades || 'A' }}
        </span>
      </div>
      <div style="display: flex; justify-content: space-between;">
        <span>Caudal:</span>
        <span :class="getStatusClass('caudal')" style="font-weight: bold;">
          {{ dispositivo.valor?.caudal || '---' }} {{ dispositivo.caudal?.unidades || 'm³/min' }}
        </span>
      </div>
    </div>
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
    },
    
    getStatusClass(tipo) {
      if (!this.dispositivo.valor || !this.dispositivo[tipo]) return '';
      
      const valor = parseFloat(this.dispositivo.valor[tipo]);
      const min = this.dispositivo[tipo].min;
      const max = this.dispositivo[tipo].max;
      const nominal = this.dispositivo[tipo].nominal;
      
      // Definir rangos de alerta
      const alertaMin = nominal * 0.9;
      const alertaMax = nominal * 1.1;
      
      if (valor < min || valor > max) {
        return 'status-danger'; // Fuera de rango - rojo
      } else if (valor < alertaMin || valor > alertaMax) {
        return 'status-warning'; // Fuera de rango nominal - amarillo
      } else {
        return 'status-success'; // Normal - verde
      }
    }
  }
}
</script>

<style scoped>
.status-success {
  color: #28a745;
}

.status-warning {
  color: #ffc107;
}

.status-danger {
  color: #dc3545;
}
</style>