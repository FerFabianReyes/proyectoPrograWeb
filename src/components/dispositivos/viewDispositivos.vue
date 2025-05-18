<template>
    <div class="card">
        <div class="card-body">
            <h3>Lista de dispositivo</h3>
            <dispositivo v-for="(item, index) in listaDispositivos" :dispositivo="item"></dispositivo>
        </div>
    </div>   
</template>

<script>
    import dispositivo from './dispositivo.vue'
    export default {
        name: 'viewDispositivos',
        props:{},
        components:{dispositivo},
        data(){
            return {listaDispositivos:[
            {id: 0,
            nombre: 'dispositivo1',
            ubicacion: 'represa tal',
            coordenadas: '20.00',
            potencia: {nominal: 7.5, min: 6.5, max: 8.5, unidades: 'kw'},
            voltaje: {},
            corriente: {},
            caudal: {}
            },
            {
            id: 0,
            nombre: 'dispositivo1',
            ubicacion: 'represa tal',
            coordenadas: '20.00',
            potencia: {nominal: 7.5, min: 6.5, max: 8.5, unidades: 'kw'},
            voltaje: {},
            corriente: {},
            caudal: {}
            }
            ]}
        }
    }
</script>