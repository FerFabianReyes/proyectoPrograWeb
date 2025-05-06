<template>
    <div class="body">
        <nav class="navbar navbar-expand-lg bg-body-tertiary">
            <div class="container-fluid">
                <a class="navbar-brand" href="#">Gatagua</a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                    data-bs-target="#navbarNavAltMarkup" aria-controls="navbarNavAltMarkup" aria-expanded="false"
                    aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarNavAltMarkup">
                    <div class="navbar-nav">
                        <a class="nav-link" aria-current="page" href="#">Agregar Bomba</a>
                        <a class="nav-link" href="#">Eliminar Bomba</a>
                        <a class="nav-link" href="#">Cerrar Sesión</a>
                    </div>
                </div>
            </div>
        </nav>
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
            estado: 'estable',
            potencia: {nominal: 7.5, min: 6.5, max: 8.5, unidades: 'kw'},
            voltaje: {},
            corriente: {},
            caudal: {}
        }
    }
}
</script>

<style></style>