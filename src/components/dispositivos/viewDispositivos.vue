<template>
    <div class="card">
        <div class="card-body">
            <h3>Lista de dispositivo</h3>
            <dispositivo v-for="(item, index) in listaDispositivos" :dispositivo="item"></dispositivo>
        </div>
    </div>   



     <div class="informacion">
            <div class="list-group">
                <a href="#" class="list-group-item list-group-item-action active" aria-current="true">
                    <div class="d-flex w-100 justify-content-between">
                        <h5 class="mb-1">{{ nombre }}</h5>
                        <small>{{ estado }}</small>
                    </div>
                    <small>{{ id }}  {{ ubicacion }}</small>
                </a>
                <a href="#" class="list-group-item list-group-item-action">
                    <div class="d-flex w-100 justify-content-between">
                        <h5 class="mb-1">Bomba 2</h5>
                        <small>Estadox</small>
                    </div>
                    <small>id y ubicación</small>
                </a>
                <a href="#" class="list-group-item list-group-item-action">
                    <div class="d-flex w-100 justify-content-between">
                        <h5 class="mb-1">Bomba 1</h5>
                        <small>Estadox</small>
                    </div>
                    <small>id y ubicación</small>
                </a>
            </div>
        </div>
        
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
            }
            ]}
        }
    }
</script>