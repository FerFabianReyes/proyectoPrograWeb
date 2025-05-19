<template>
    <div class="lista-view-container" style="display: flex; gap: 20px; height: 80vh;">
        <!-- Contenedor con scroll vertical -->
        <div class="lista-dispositivos"
            style="flex: 1; min-width: 400px; border-right: 1px solid #ccc; padding-right: 20px; overflow-y: auto;">
            <h1>Lista de dispositivos</h1>
            <dispositivo v-for="(item, index) in listaDispositivos" :key="index" :dispositivo="item" />
        </div>

        <div class="analisis-bomba" style="flex: 1; min-width: 750px;">
            <router-view />
        </div>
    </div>
</template>


<script>
import dispositivo from './dispositivo.vue'

export default {
    name: 'listaView',
    components: { dispositivo },
    data() {
        return {
            listaDispositivos: []
        }
    },
    mounted() {
        const bombas = JSON.parse(localStorage.getItem('bombas')) || [];
        this.listaDispositivos = bombas.map(bomba => ({
            nombre: bomba.nombre,
            estado: bomba.estado,
            ubicacion: bomba.coordenadas // alias para ubicacion
        }));
    }
}
</script>
