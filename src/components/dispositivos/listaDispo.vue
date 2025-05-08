<template>
    <div class="card">
        <div class="card-body">
            <p>id: {{ dispositivo.id }}</p>
            <p>Nombre: {{ dispositivo.nombre }}</p>
            <p>Ubicación: {{ dispositivo.ubicacion }}</p>
            <p>Coordenadas: {{ dispositivo.coordenadas }}</p>
            <p>Potencia nom: {{ dispositivo.potencia.nominal }} {{ dispositivo.potencia.unidades }}</p>
            <p>Potencia min: {{ dispositivo.potencia.min }} {{ dispositivo.potencia.unidades }}</p>
            <p>Potencia max: {{ dispositivo.potencia.max }} {{ dispositivo.potencia.unidades }}</p>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Dispositivo',
    props: {
        dispositivo: {
            type: Object
        }
    },
    data() {
        return {
            dispositivo_: {
            id: 0,
            nombre: 'dispositivo1',
            ubicacion: 'represa tal',
            coordenadas: '20.00',
            potencia: {nominal: 7.5, min: 6.5, max: 8.5, unidades: 'kw'},
            voltaje: {},
            corriente: {},
            caudal: {}
        }
    }
    }
}
</script>