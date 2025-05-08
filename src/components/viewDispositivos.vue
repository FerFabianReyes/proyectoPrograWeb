<template>
    <div class="card">
        <div class="card-body">
            <h3>Lista de dispositivo</h3>
            <Dispositivo v-for="(item.index) in listaDispo"></Dispositivo>
        </div>
    </div>   
</template>

<script>
    import Dispositivo from './dispositivo.vue'
    export default {
        name: 'viewDispositivos',
        properties{
            propiedades:{
                type: Object
                require
            }
        },
        DataTransfer(){
            listaDispo:[
                

            ]
        }
    }
</script>