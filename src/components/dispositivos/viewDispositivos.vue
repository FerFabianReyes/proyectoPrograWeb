<template>
    <div class="lista-view-container" style="display: flex; gap: 20px; height: 80vh;">
        <!-- Contenedor con scroll vertical -->
        <div class="lista-dispositivos"
            style="flex: 1; min-width: 400px; border-right: 1px solid #ccc; padding-right: 20px; overflow-y: auto;">
            <h1>Lista de dispositivos</h1>
            <dispositivo v-for="(item, index) in listaDispositivos" :key="index" :dispositivo="item" /> <!-- recorre listaDispositivos y le pasa la inforamción de la bomba al hijo dispositivo-->>
        </div>

        <div class="analisis-bomba" style="flex: 1; min-width: 800px;">
            <router-view /> <!--Es para mostrar otro componente-->
        </div>
    </div>
</template>


<script>
import dispositivo from './dispositivo.vue'

export default {
    name: 'listaView',
    components: { dispositivo }, //importa el componente dispositivo
    data() {
        return {
            listaDispositivos: [] //es el arreglo de bombas que se va a mostrar
        }
    },
    mounted() { //se ejecuta cuando el componente se monta
        const usuarioActivo = JSON.parse(localStorage.getItem('usuarioActivo'));
        const bombas = usuarioActivo?.bombas || [];
        this.listaDispositivos = bombas.map(bomba => ({
            nombre: bomba.nombre,
            estado: bomba.estado,
            ubicacion: bomba.ubicacion
        }));
    }
}
</script>
