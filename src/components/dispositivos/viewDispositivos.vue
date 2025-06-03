<template>
    <div class="lista-view-container" style="display: flex; gap: 20px; height: 80vh;">
        <!-- Contenedor con scroll vertical -->
        <div class="lista-dispositivos"
            style="flex: 1; min-width: 400px; border-right: 1px solid #ccc; padding-right: 20px; overflow-y: auto;">
            <h1>Lista de dispositivos</h1>
            <dispositivo v-for="(item, index) in listaDispositivos" :key="index" :dispositivo="item"
                @seleccionar="seleccionarBomba" />
        </div>

        <div class="analisis-bomba" style="flex: 1; min-width: 800px;">

            <div class="dropdown">
                <button class="btn dropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
                    Monitoreo
                </button>
                <ul class="dropdown-menu">
                    <li><a class="dropdown-item" @click="refresh">Refresh</a></li>
                    <li><a class="dropdown-item" @click="start">Start</a></li>
                    <li><a class="dropdown-item" @click="stop">Stop</a></li>
                </ul>
            </div>

            <analisis-bomba v-if="bombaSeleccionada" :bomba="bombaSeleccionada" />
        </div>
    </div>
</template>

<script>
import dispositivo from './dispositivo.vue'
import analisisBomba from './analisisBomba.vue'

export default {
    name: 'listaView',
    components: { dispositivo, analisisBomba },
    data() {
        return {
            listaDispositivos: [],
            bombaSeleccionada: null,
            nIntervId: null
        }
    },
    mounted() {
        const usuarioActivo = JSON.parse(localStorage.getItem('usuarioActivo'));
        const bombas = usuarioActivo?.bombas || [];
        this.listaDispositivos = bombas;

        // Inicializar valores si no existen
        this.listaDispositivos.forEach((bomba, index) => {
            if (!bomba.valor) {
                bomba.valor = {
                    potencia: '0.000',
                    voltaje: '0.0',
                    corriente: '0.00',
                    caudal: '0.000'
                };
            }
        });
    },
    beforeUnmount() {
        if (this.nIntervId) {
            clearInterval(this.nIntervId);
        }
    },
    methods: {
        seleccionarBomba(bomba) {
            this.bombaSeleccionada = bomba;
        },

        refresh() {

            this.listaDispositivos.forEach((item, index) => {
                const potenciaMin = parseFloat(item.potencia.min);
                const potenciaMax = parseFloat(item.potencia.max);
                const voltajeMin = parseFloat(item.voltaje.min);
                const voltajeMax = parseFloat(item.voltaje.max);
                const corrienteMin = parseFloat(item.corriente.min);
                const corrienteMax = parseFloat(item.corriente.max);
                const caudalMin = parseFloat(item.caudal.min);
                const caudalMax = parseFloat(item.caudal.max);

                let potencia = potenciaMin + ((potenciaMax - potenciaMin) * Math.random());
                let voltaje = voltajeMin + ((voltajeMax - voltajeMin) * Math.random());
                let corriente = corrienteMin + ((corrienteMax - corrienteMin) * Math.random());
                let caudal = caudalMin + ((caudalMax - caudalMin) * Math.random());

                // Crear objeto con nuevos valores
                const nuevosValores = {
                    potencia: potencia.toFixed(3),
                    voltaje: voltaje.toFixed(1),
                    corriente: corriente.toFixed(2),
                    caudal: caudal.toFixed(3)
                };

                this.listaDispositivos[index].valor = nuevosValores;

                // Actualizar bomba seleccionada si coincide
                if (this.bombaSeleccionada && this.bombaSeleccionada.nombre === item.nombre) {
                    this.bombaSeleccionada = { ...this.listaDispositivos[index] };
                }
            });

            // Forzar actualización de la vista
            this.$forceUpdate();
        },

        start() {
            if (!this.nIntervId) {
                this.nIntervId = setInterval(() => {
                    this.refresh();
                }, 1000);
                console.log('Intervalo iniciado:', this.nIntervId);
            }
        },

        stop() {
            if (this.nIntervId) {
                clearInterval(this.nIntervId);
                this.nIntervId = null;
                console.log('Intervalo detenido');
            }
        }
    }
}
</script>