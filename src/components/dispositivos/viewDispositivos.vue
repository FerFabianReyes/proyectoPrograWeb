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
                <button class="btn dropdown-toggle" type="button" data-bs-toggle="dropdown"
                    aria-expanded="false">
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
            bombaSeleccionada: null
        }
    },
    mounted() {
        const usuarioActivo = JSON.parse(localStorage.getItem('usuarioActivo'));
        const bombas = usuarioActivo?.bombas || [];
        this.listaDispositivos = bombas;
    },
    methods: {
        seleccionarBomba(bomba) {
            this.bombaSeleccionada = bomba;
        },

        refresh() {
                this.dispositivos.forEach((item,index)=>{
                    let desviacion=0.05

                    let potencia = item.potencia.min +
                                (((item.potencia.max - item.potencia.min) +
                                (item.potencia.nominal * 0.10)) * Math.random()) 
                    let voltaje = item.voltaje.min +
                                (((item.voltaje.max - item.voltaje.min) +
                                (item.voltaje.nominal * desviacion)) * Math.random())
                    let corriente = item.corriente.min +
                                (((item.corriente.max - item.corriente.min) +
                                (item.corriente.nominal * desviacion)) * Math.random())
                    let caudal = item.caudal.min +
                                (((item.caudal.max - item.caudal.min) +
                                (item.caudal.nominal * 0.10)) * Math.random())           
                    
                    this.dispositivos[index].opera = {
                        potencia: potencia,
                        voltaje: voltaje,
                        corriente: corriente,
                        caudal: caudal
                    }
                })                
            },
            start() {
                if (!this.nIntervId) {
                    this.nIntervId = setInterval(()=>{this.refresh()}, 1000);
                }
            },
            stop() {
                clearInterval(this.nIntervId);
                this.nIntervId = null;
            }

    }
}
</script>
