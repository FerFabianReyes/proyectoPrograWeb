<template>
        <h1>Lista de dispositivos</h1> <br>
            <dispositivo v-for="(item, index) in listaDispositivos" :dispositivo="item"></dispositivo>

        
        <table class="table table-striped table-hover">
            <tr>
                <th>Propiedades</th>
                <th>x</th>
                <th>x</th>
                <th>x</th>
            </tr>
            <tr>
                <td>Voltaje</td>
                <td></td>
                <td></td>
                <td></td>
            </tr>
            <tr>
                <td>Corriente</td>
                <td></td>
                <td></td>
                <td></td>
            </tr>
            <tr>
                <td>Caudal</td>
                <td></td>
                <td></td>
                <td></td>
            </tr>
        </table>

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
            id: 1,
            nombre: 'dispositivo2',
            ubicacion: 'represa tal 2',
            coordenadas: '20.00',
            potencia: {nominal: 7.5, min: 6.5, max: 8.5, unidades: 'kw'},
            voltaje: {},
            corriente: {},
            caudal: {}
            },
            {
            id: 2,
            nombre: 'dispositivo3',
            ubicacion: 'represa tal 3',
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