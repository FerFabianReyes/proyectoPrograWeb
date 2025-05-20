<template>
    <div class="analisis"
        style="flex: 1; min-width: 800px; border-right: 1px solid #ccc; padding-right: 20px; overflow-y: auto; max-height: 80vh;">

        <div class="container text-center">
            <div class="row">
                <div class="col">
                    <p class="h4" v-if="bomba">{{ bomba.nombre }}</p>
                </div>
                <div class="col">
                    <p class="h4" v-if="bomba">{{ bomba.estado }}</p>
                </div>
                <div class="col">
                    <button type="button" class="btn" @click="recargarPagina">
                        <span class="material-symbols-outlined">refresh</span>
                    </button>

                </div>
            </div>
            <br>
            <div class="row">
                <div class="col">
                    <p class="h5" v-if="bomba">Ubicación</p>
                    <p class="h6" v-if="bomba">{{ bomba.ubicacion }}</p>
                </div>
                <div class="col">
                    <p class="h5" v-if="bomba">Coordenadas</p>
                    <p class="h6" v-if="bomba">{{ bomba.coordenadas }}</p>
                </div>
                <div class="col">
                    <p class="h5" v-if="bomba">Fecha de registro</p>
                    <p class="h6" v-if="bomba">{{ bomba.fechaRegistro }}</p>
                </div>
            </div>
        </div>

        <div class="accordion" id="accordionPanelsStayOpenExample">
            <div class="accordion-item">
                <h2 class="accordion-header">
                    <button class="accordion-button" type="button" data-bs-toggle="collapse"
                        data-bs-target="#panelsStayOpen-collapseOne" aria-expanded="false"
                        aria-controls="panelsStayOpen-collapseOne">
                        Potencia
                    </button>
                </h2>
                <div id="panelsStayOpen-collapseOne" class="accordion-collapse collapse show">
                    <div class="accordion-body">
                        <table class="table table-hover" v-if="bomba">
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
                        <table class="table table-hover" v-if="bomba">
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
            <div class="accordion-item">
                <h2 class="accordion-header">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse"
                        data-bs-target="#panelsStayOpen-collapseThree" aria-expanded="false"
                        aria-controls="panelsStayOpen-collapseThree">
                        Corriente
                    </button>
                </h2>
                <div id="panelsStayOpen-collapseThree" class="accordion-collapse collapse">
                    <div class="accordion-body">
                        <table class="table table-hover" v-if="bomba">
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
            <div class="accordion-item">
                <h2 class="accordion-header">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse"
                        data-bs-target="#panelsStayOpen-collapseFour" aria-expanded="false"
                        aria-controls="panelsStayOpen-collapseThree">
                        Caudal
                    </button>
                </h2>
                <div id="panelsStayOpen-collapseFour" class="accordion-collapse collapse">
                    <div class="accordion-body">
                        <table class="table table-hover" v-if="bomba">
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



    </div>

</template>

<script>
export default {
    name: 'analisisBomba',
    data() {
        return {
            bomba: null,
            mostrarModal: false,
        };
    },
    mounted() {
        this.cargarBomba();
    },
    watch: {
        '$route.query.nombre'() {
            this.cargarBomba();
        }
    },
    methods: {
        cargarBomba() {
            const nombre = this.$route.query.nombre;
            const bombas = JSON.parse(localStorage.getItem('bombas')) || [];
            this.bomba = bombas.find(b => b.nombre === nombre) || null;

            if (this.bomba) {
                this.cambiarEstadoBomba();
            }
        },
        cambiarEstadoBomba() {
            const r = Math.random();
            if (r < 1 / 6) {
                this.bomba.estado = "Crítico";
            } else if (r < 4 / 6) {
                this.bomba.estado = "Advertencia";
            } else {
                this.bomba.estado = "Estable";
            }

            const bombas = JSON.parse(localStorage.getItem('bombas')) || [];
            const index = bombas.findIndex(b => b.nombre === this.bomba.nombre);
            if (index !== -1) {
                bombas[index].estado = this.bomba.estado;
                localStorage.setItem('bombas', JSON.stringify(bombas));
            }
        },
        recargarPagina() {
            window.location.reload();
        },
        verifEliminar() {
            this.mostrarModal = true;
        },
        eliminarBomba() {
            if (!this.bomba) return;

            const bombas = JSON.parse(localStorage.getItem('bombas')) || [];
            const index = bombas.findIndex(b => b.nombre === this.bomba.nombre);
            if (index !== -1) {
                bombas.splice(index, 1);
                localStorage.setItem('bombas', JSON.stringify(bombas));
            }

            this.mostrarModal = false;
            this.$router.push('/menu').then(() => {
                window.location.reload();
            });

        },
        cerrarModal() {
            this.mostrarModal = false;
        },
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