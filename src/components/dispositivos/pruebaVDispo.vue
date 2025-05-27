<template>
<div class="card">
    <div class="btn-toolbar" role="toolbar" aria-label="Toolbar with button groups">
    <div class="btn-group me-2" role="group" aria-label="Second group">
        <button type="button" class="btn btn-secondary" @click="refresh">refresh</button>
        <button type="button" class="btn btn-secondary" @click="start">Iniciar</button>
        <button type="button" class="btn btn-secondary" @click="stop">Detener</button>
    </div>
    </div>
    <div class="card-body">
        <Dispositivo  v-for="(item,index) in dispositivos" :key="index" :identifica="item.identifica" :opera="item.opera"></Dispositivo>
    </div>
</div>
</template>
<script>
   // import DSDispositivos from '@/servicios/dispositivos'
    import Dispositivo from './pruebaDispositivo.vue'

    export default {
        name: 'ViewDispositivo',
        components: {
            Dispositivo
        },
        data() {
            return {
                nIntervId: null,
                dispositivos: [
                    {
                        identifica: {
                            identificador: 1,
                            nombre: 'Represa Uno',
                            ubicacion: 'En la compuerta de la Represa Uno',
                            coordenadas: '19.7060° N, 101.1950° W',
                            idestatus: 1,
                            estatus: 'Operacion Normal',
                            potencia:  { nominal: 7.400, minimo: 6.200, maximo: 8.600, um: 'KW' },
                            voltaje: { nominal: 240, minimo: 230, maximo: 250, um: 'Volts' },
                            corriente: { nominal: 30, minimo: 25, maximo: 35, um: 'Amperes' },
                            caudal: { nominal: 1, minimo: 0.10, maximo: 1.20, um: 'm3/minuto' },
                            fechaRegistro: new Date().toUTCString()
                        },
                        opera: {
                            voltaje: 240,
                            corriente: 30,
                            caudal: 1,
                            estatus: 1    // {1->Normal, 2->Advertencia, 3->Error}
                        },
                        estado: 1   // {1->Encendido, 2->Apagado, 3->Bloqueado}
                    },
                    {
                        identifica: {
                            identificador: 2,
                            nombre: 'Represa dos',
                            ubicacion: 'En la compuerta de la Represa dos',
                            coordenadas: '19.7060° N, 101.1950° W',
                            idestatus: 1,
                            estatus: 'Operacion Normal',
                            potencia:  { nominal: 7.400, minimo: 6.200, maximo: 8.600, um: 'KW' },
                            voltaje: { nominal: 240, minimo: 230, maximo: 250, um: 'Volts' },
                            corriente: { nominal: 30, minimo: 25, maximo: 35, um: 'Amperes' },
                            caudal: { nominal: 1, minimo: 0.10, maximo: 1.20, um: 'm3/minuto' },
                            fechaRegistro: new Date().toUTCString()
                        },
                        opera: {
                            voltaje: 240,
                            corriente: 30,
                            caudal: 1,
                            estatus: 1    // {1->Normal, 2->Advertencia, 3->Error}
                        },
                        estado: 1   // {1->Encendido, 2->Apagado, 3->Bloqueado}
                    },
                    {
                        identifica: {
                            identificador: 3,
                            nombre: 'Represa tres',
                            ubicacion: 'En la compuerta de la Represa tres',
                            coordenadas: '19.7060° N, 101.1950° W',
                            idestatus: 1,
                            estatus: 'Operacion Normal',
                            potencia:  { nominal: 7.400, minimo: 6.200, maximo: 8.600, um: 'KW' },
                            voltaje: { nominal: 240, minimo: 230, maximo: 250, um: 'Volts' },
                            corriente: { nominal: 30, minimo: 25, maximo: 35, um: 'Amperes' },
                            caudal: { nominal: 1, minimo: 0.10, maximo: 1.20, um: 'm3/minuto' },
                            fechaRegistro: new Date().toUTCString()
                        },
                        opera: {
                            voltaje: 240,
                            corriente: 30,
                            caudal: 1,
                            estatus: 1    // {1->Normal, 2->Advertencia, 3->Error}
                        },
                        estado: 1   // {1->Encendido, 2->Apagado, 3->Bloqueado}
                    },
                    {
                        identifica: {
                            identificador: 4,
                            nombre: 'Represa cuatro',
                            ubicacion: 'En la compuerta de la Represa Cuatro',
                            coordenadas: '19.7060° N, 101.1950° W',
                            idestatus: 1,
                            estatus: 'Operacion Normal',
                            potencia:  { nominal: 7.400, minimo: 6.200, maximo: 8.600, um: 'KW' },
                            voltaje: { nominal: 240, minimo: 230, maximo: 250, um: 'Volts' },
                            corriente: { nominal: 30, minimo: 25, maximo: 35, um: 'Amperes' },
                            caudal: { nominal: 1, minimo: 0.10, maximo: 1.20, um: 'm3/minuto' },
                            fechaRegistro: new Date().toUTCString()
                        },
                        opera: {
                            voltaje: 240,
                            corriente: 30,
                            caudal: 1,
                            estatus: 1    // {1->Normal, 2->Advertencia, 3->Error}
                        },
                        estado: 1   // {1->Encendido, 2->Apagado, 3->Bloqueado}
                    },
                    {
                        identifica: {
                            identificador: 5,
                            nombre: 'Represa Cinco',
                            ubicacion: 'En la compuerta de la Represa Cinco',
                            coordenadas: '19.7060° N, 101.1950° W',
                            idestatus: 1,
                            estatus: 'Operacion Normal',
                            potencia:  { nominal: 7.400, minimo: 6.200, maximo: 8.600, um: 'KW' },
                            voltaje: { nominal: 240, minimo: 230, maximo: 250, um: 'Volts' },
                            corriente: { nominal: 30, minimo: 25, maximo: 35, um: 'Amperes' },
                            caudal: { nominal: 1, minimo: 0.10, maximo: 1.20, um: 'm3/minuto' },
                            fechaRegistro: new Date().toUTCString()
                        },
                        opera: {
                            voltaje: 240,
                            corriente: 30,
                            caudal: 1,
                            estatus: 1    // {1->Normal, 2->Advertencia, 3->Error}
                        },
                        estado: 1   // {1->Encendido, 2->Apagado, 3->Bloqueado}
                    }
                ]
            }
        },
        methods:{
            refresh() {
                this.dispositivos.forEach((item,index)=>{
                    let desviacion=0.05
                    let voltaje = item.identifica.voltaje.minimo +
                                (((item.identifica.voltaje.maximo - item.identifica.voltaje.minimo) +
                                (item.identifica.voltaje.nominal * desviacion)) * Math.random())
                    let corriente = item.identifica.corriente.minimo +
                                (((item.identifica.corriente.maximo - item.identifica.corriente.minimo) +
                                (item.identifica.corriente.nominal * desviacion)) * Math.random())
                    let caudal = item.identifica.caudal.minimo +
                                (((item.identifica.caudal.maximo - item.identifica.caudal.minimo) +
                                (item.identifica.caudal.nominal * 0.10)) * Math.random())
                    
                    this.dispositivos[index].opera = {
                        voltaje: voltaje,
                        corriente: corriente,
                        caudal: caudal,
                        estatus: 1
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
            },
            nuevo(dispositivo) {
                this.dispositivos.push(disppsitovo)
            }
        },
        created () {
            //DSDispositivos.init()
            //this.objDispositivos = DSDispositivos.getListaDataStore()
        }
    }
    </script>
    