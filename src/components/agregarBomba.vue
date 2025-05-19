<template>
  <div>
    <p class="h1">Agregar Bomba</p> <br><br>
    <p class="h4">Ingrese la información de la bomba en los campos</p> <br>

    <div class="mb-2">
      <input v-model="bomba.nombre" type="text" class="form-control" placeholder="Nombre">
      <input v-model="bomba.estado" type="text" class="form-control" placeholder="Estado">
      <input v-model="bomba.ubicacion" type="text" class="form-control" placeholder="Ubicación">
      <input v-model="bomba.coordenadas" type="text" class="form-control" placeholder="Coordenadas">
    </div>

    <div class="mb-2">
      <label class="h6">Potencia</label>
      <input v-model="bomba.potencia.nominal" type="text" class="form-control" placeholder="Nominal">
      <input v-model="bomba.potencia.min" type="text" class="form-control" placeholder="Mínima">
      <input v-model="bomba.potencia.max" type="text" class="form-control" placeholder="Máxima">
      <input v-model="bomba.potencia.unidades" type="text" class="form-control" placeholder="Unidades">
    </div>

    <div class="mb-2">
      <label class="h6">Voltaje</label>
      <input v-model="bomba.voltaje.nominal" type="text" class="form-control" placeholder="Nominal">
      <input v-model="bomba.voltaje.min" type="text" class="form-control" placeholder="Mínima">
      <input v-model="bomba.voltaje.max" type="text" class="form-control" placeholder="Máxima">
      <input v-model="bomba.voltaje.unidades" type="text" class="form-control" placeholder="Unidades">
    </div>

    <div class="mb-2">
      <label class="h6">Corriente</label>
      <input v-model="bomba.corriente.nominal" type="text" class="form-control" placeholder="Nominal">
      <input v-model="bomba.corriente.min" type="text" class="form-control" placeholder="Mínima">
      <input v-model="bomba.corriente.max" type="text" class="form-control" placeholder="Máxima">
      <input v-model="bomba.corriente.unidades" type="text" class="form-control" placeholder="Unidades">
    </div>

    <div class="mb-2">
      <label class="h6">Caudal</label>
      <input v-model="bomba.caudal.nominal" type="text" class="form-control" placeholder="Nominal">
      <input v-model="bomba.caudal.min" type="text" class="form-control" placeholder="Mínima">
      <input v-model="bomba.caudal.max" type="text" class="form-control" placeholder="Máxima">
      <input v-model="bomba.caudal.unidades" type="text" class="form-control" placeholder="Unidades">
    </div>

    <button @click="guardarBomba" class="btn btn-outline-primary">Agregar</button>
  </div>
</template>

<script>
export default {
  name: 'AgregarBomba',
  data() {
    return {
      bomba: {
        id: Date.now(),
        nombre: '',
        estado: '',
        ubicacion: '',
        coordenadas: '',
        fechaRegistro: new Date().toISOString(),
        potencia: { nominal: '', min: '', max: '', unidades: '' },
        voltaje: { nominal: '', min: '', max: '', unidades: '' },
        corriente: { nominal: '', min: '', max: '', unidades: '' },
        caudal: { nominal: '', min: '', max: '', unidades: '' },
      }
    }
  },
  methods: {
  camposLlenos(obj) {
    // Verifica que todos los valores del objeto no estén vacíos
    return Object.values(obj).every(v => v !== '');
  },
  guardarBomba() {
    const b = this.bomba;

    const todosLlenos =
      b.nombre && b.estado && b.coordenadas &&
      this.camposLlenos(b.potencia) &&
      this.camposLlenos(b.voltaje) &&
      this.camposLlenos(b.corriente) &&
      this.camposLlenos(b.caudal);

    if (!todosLlenos) {
      alert('Por favor, complete todos los campos antes de guardar.');
      return;
    }

    let bombas = JSON.parse(localStorage.getItem('bombas')) || [];
    bombas.push(this.bomba);
    localStorage.setItem('bombas', JSON.stringify(bombas));
    alert('Bomba guardada');
    this.$router.push('/menu');
  }
}

}
</script>

<style>
.mb-2 {
  width: 400px;
  padding: 20px;
}
</style>
