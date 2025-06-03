<template>
  <div class="analisis"
       style="flex: 1; min-width: 800px; border-right: 1px solid #ccc; padding-right: 20px; overflow-y: auto; max-height: 80vh;">
    
    <div v-if="bomba">
      <div class="container text-center">
        <div class="row">
          <div class="col">
            <p class="h4">{{ bomba.nombre }}</p>
          </div>
          <div class="col">
            <p class="h4">{{ bomba.estado || 'Activo' }}</p>
          </div>
        </div>
        <br>
        <div class="row">
          <div class="col">
            <p class="h5">Ubicación</p>
            <p class="h6">{{ bomba.ubicacion }}</p>
          </div>
          <div class="col">
            <p class="h5">Coordenadas</p>
            <p class="h6">{{ bomba.coordenadas || 'N/A' }}</p>
          </div>
          <div class="col">
            <p class="h5">Fecha de registro</p>
            <p class="h6">{{ bomba.fechaRegistro || 'N/A' }}</p>
          </div>
        </div>
      </div>

      <div class="accordion" id="accordionPanelsStayOpenExample">
        <!-- Potencia -->
        <div class="accordion-item">
          <h2 class="accordion-header">
            <button class="accordion-button" type="button" data-bs-toggle="collapse"
                    data-bs-target="#panelsStayOpen-collapseOne" aria-expanded="true"
                    aria-controls="panelsStayOpen-collapseOne">
              Potencia
            </button>
          </h2>
          <div id="panelsStayOpen-collapseOne" class="accordion-collapse collapse show">
            <div class="accordion-body">
              <table class="table table-hover">
                <tr>
                  <th>Nominal</th>
                  <th>Máxima</th>
                  <th>Mínima</th>
                  <th>Unidades</th>
                </tr>
                <tr>
                  <td>{{ bomba.potencia.nominal }}</td>
                  <td>{{ bomba.potencia.max }}</td>
                  <td>{{ bomba.potencia.min }}</td>
                  <td>{{ bomba.potencia.unidades }}</td>
                </tr>
              </table>
            </div>
          </div>
        </div>

        <!-- Voltaje -->
        <div class="accordion-item">
          <h2 class="accordion-header">
            <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse"
                    data-bs-target="#panelsStayOpen-collapseTwo" aria-expanded="false"
                    aria-controls="panelsStayOpen-collapseTwo">
              Voltaje
            </button>
          </h2>
          <div id="panelsStayOpen-collapseTwo" class="accordion-collapse collapse">
            <div class="accordion-body">
             <table class="table table-hover">
                <tr>
                  <th>Nominal</th>
                  <th>Máxima</th>
                  <th>Mínima</th>
                  <th>Unidades</th>
                </tr>
                <tr>
                  <td>{{ bomba.voltaje.nominal }}</td>
                  <td>{{ bomba.voltaje.max }}</td>
                  <td>{{ bomba.voltaje.min }}</td>
                  <td>{{ bomba.voltaje.unidades }}</td>
                </tr>
              </table>
            </div>
          </div>
        </div>

        <!-- Corriente -->
        <div class="accordion-item" v-if="bomba.potencia">
          <h2 class="accordion-header">
            <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse"
                    data-bs-target="#panelsStayOpen-collapseThree" aria-expanded="false"
                    aria-controls="panelsStayOpen-collapseThree">
              Corriente
            </button>
          </h2>
          <div id="panelsStayOpen-collapseThree" class="accordion-collapse collapse">
            <div class="accordion-body">
              <table class="table table-hover">
                <tr>
                  <th>Nominal</th>
                  <th>Máxima</th>
                  <th>Mínima</th>
                  <th>Unidades</th>
                </tr>
                <tr>
                  <td>{{ bomba.corriente.nominal }}</td>
                  <td>{{ bomba.corriente.max }}</td>
                  <td>{{ bomba.corriente.min }}</td>
                  <td>{{ bomba.corriente.unidades }}</td>
                </tr>
              </table>
            </div>
          </div>
        </div>

        <!-- Caudal -->
        <div class="accordion-item" v-if="bomba.voltaje">
          <h2 class="accordion-header">
            <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse"
                    data-bs-target="#panelsStayOpen-collapseFour" aria-expanded="false"
                    aria-controls="panelsStayOpen-collapseFour">
              Caudal
            </button>
          </h2>
          <div id="panelsStayOpen-collapseFour" class="accordion-collapse collapse">
            <div class="accordion-body">
              <table class="table table-hover">
                <tr>
                  <th>Nominal</th>
                  <th>Máxima</th>
                  <th>Mínima</th>
                  <th>Unidades</th>
                </tr>
                <tr>
                  <td>{{ bomba.caudal.nominal }}</td>
                  <td>{{ bomba.caudal.max }}</td>
                  <td>{{ bomba.caudal.min }}</td>
                  <td>{{ bomba.caudal.unidades }}</td>
                </tr>
              </table>
            </div>
          </div>
        </div>
      </div>

      <br>
      <button type="button" class="btn btn-danger" @click="verifEliminar">Eliminar bomba</button>

      <!-- MODAL -->
      <div class="modal fade" :class="{ 'show d-block': mostrarModal }" tabindex="-1" v-if="mostrarModal">
        <div class="modal-dialog modal-dialog-centered">
          <div class="modal-content p-3">
            <div class="modal-header">
              <h5 class="modal-title">Eliminar Bomba</h5>
              <button type="button" class="btn-close" @click="cerrarModal"></button>
            </div>
            <div class="modal-body">
              <p>¿Estás seguro que deseas eliminar esta bomba?</p>
            </div>
            <div class="modal-footer">
              <button class="btn btn-secondary" @click="cerrarModal">Cerrar</button>
              <button class="btn btn-danger" @click="eliminarBomba">Eliminar</button>
            </div>
          </div>
        </div>
      </div>
      <!-- FIN MODAL -->

    </div>
    
    <div v-else>
      <div class="container text-center">
        <p class="h4">No se encontró la bomba seleccionada.</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AnalisisBomba',
  props: {
    bomba: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      mostrarModal: false,
    };
  },
  methods: {
    eliminarBomba() {
      const usuarioActivo = localStorage.getItem('usuarioActivo');
      const bombasKey = `bombas_${usuarioActivo}`;
      const bombasGuardadas = JSON.parse(localStorage.getItem(bombasKey)) || [];

      const bombasActualizadas = bombasGuardadas.filter(b => b.nombre !== this.bomba.nombre);
      localStorage.setItem(bombasKey, JSON.stringify(bombasActualizadas));

      this.mostrarModal = false;
      // Redirigir y recargar para reflejar los cambios
      this.$router.push('/menu').then(() => {
        window.location.reload();
      });
    },
    recargarPagina() {
      window.location.reload();
    },
    verifEliminar() {
      this.mostrarModal = true;
    },
    cerrarModal() {
      this.mostrarModal = false;
    },
  }
};
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

.analisis {
  padding: 20px;
}

.accordion-button:not(.collapsed) {
  background-color: #e3f2fd;
  border-color: #90caf9;
}

.table th {
  background-color: #f8f9fa;
  font-weight: 600;
}

.btn {
  margin: 5px;
}
</style>