<template>
    <div class="card">
        <div class="card-body">
            <p>id: {{ id }}</p>
            <p>Nombre: {{ nombre }}</p>
            <p>Ubicación: {{ ubicacion }}</p>
            <p>Coordenadas: {{ coordenadas }}</p>
            <p>Potencia nom: {{ potencia.nominal }} {{ potencia.unidades }}</p>
            <p>Potencia min: {{ potencia.min }} {{ potencia.unidades }}</p>
            <p>Potencia max: {{ potencia.max }} {{ potencia.unidades }}</p>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Dispositivo',
    data() {
        return {
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
</script>