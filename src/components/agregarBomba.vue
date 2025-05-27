<template>
  <div>
    <p class="h1">Agregar Bomba</p> <br><br>
    <p class="h4">Ingrese la información de la bomba en los campos</p> <br>

    <div class="mb-2">
      <input v-model="bomba.nombre" type="text" class="form-control" placeholder="Nombre">
      <input v-model="bomba.estado" type="text" class="form-control" placeholder="Estado">
      <input v-model="bomba.nombre" type="text" class="form-control" placeholder="Nombre"> <!--Si se ingresa algo. v-model se cambia el valor a la bomba con lo que le puso-->
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

    <button @click="verificarCampos" class="btn btn-outline-primary">Agregar</button>

    <!-- MODAL -->
    <div class="modal fade show d-block" tabindex="-1" v-if="mostrarModal">
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content p-3">
          <div class="modal-header">
            <h5 class="modal-title">Confirmar guardado</h5>
            <button type="button" class="btn-close" @click="cerrarModal"></button>
          </div>
          <div class="modal-body">
            <p>¿Estás seguro que deseas agregar esta bomba?</p>
          </div>
          <div class="modal-footer">
            <button class="btn btn-secondary" @click="cerrarModal">Cerrar</button>
            <button class="btn btn-primary" @click="confirmarGuardarBomba">Guardar</button>
          </div>
        </div>
      </div>
    </div>
    <!-- FIN MODAL -->

  </div>
</template>

<script>
export default {
  name: 'AgregarBomba',
  data() {
    return {
      mostrarModal: false,
      bomba: {
        id: Date.now(),
        nombre: '',
        estado: '',
        estado: '',
        ubicacion: '',
        coordenadas: '',
        fechaRegistro: new Date().toISOString(),
        fechaRegistro: new Date().toISOString(),
        potencia: { nominal: '', min: '', max: '', unidades: '' },
        voltaje: { nominal: '', min: '', max: '', unidades: '' },
        corriente: { nominal: '', min: '', max: '', unidades: '' },
        caudal: { nominal: '', min: '', max: '', unidades: '' },
        valores: { pot_valor:'', vol_valor:'', corr_valor:'', caud_valor:'' }
      }
      }
    }
  },
  methods: {
    camposLlenos(obj) {
      return Object.values(obj).every(v => v !== '');
    },
    verificarCampos() {
      const b = this.bomba;

      const todosLlenos =
        b.nombre && b.estado && b.coordenadas &&
        b.nombre && b.estado && b.coordenadas &&
        this.camposLlenos(b.potencia) &&
        this.camposLlenos(b.voltaje) &&
        this.camposLlenos(b.corriente) &&
        this.camposLlenos(b.caudal);

      if (!todosLlenos) {
        alert('Por favor, complete todos los campos antes de guardar.');
        return;
      }

      this.mostrarModal = true;
    },
    cerrarModal() {
      this.mostrarModal = false;
    },
    confirmarGuardarBomba() {
      // carga el usuario activo
      let usuarioActivo = JSON.parse(localStorage.getItem('usuarioActivo'));
      let usuarios = JSON.parse(localStorage.getItem('usuarios')) || [];

      if (!usuarioActivo.bombas) {
        usuarioActivo.bombas = [];
      }

      // vrear una copia de la bomba con un nuevo id y fecha
      const nuevaBomba = {
        ...this.bomba,
        id: Date.now(),
        fechaRegistro: new Date().toISOString()
      };

      // agregar la bomba al usuario
      usuarioActivo.bombas.push(nuevaBomba);

      // actualizar el usuario en el array de usuarios
      usuarios = usuarios.map(u =>
        u.email === usuarioActivo.email ? usuarioActivo : u
      );

      // guardar actualizaciones
      localStorage.setItem('usuarios', JSON.stringify(usuarios));
      localStorage.setItem('usuarioActivo', JSON.stringify(usuarioActivo));

      // tedirigir o cerrar modal
      this.mostrarModal = false;
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

.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 1050;
}

.modal-content {
  background-color: #fff;
  border-radius: 10px;
}
</style>
