<template>
<div class="col">
    <div class="card">
        <img src="..." class="card-img-top" alt="...">
        <div class="card-body">
            <h5 class="card-title">{{identifica.nombre}}</h5>
            <p class="card-text">{{ identifica.ubicacion }}, {{ identifica.coordenadas }}.</p>
            <table class="table table-bordered">
                <thead>
                    <tr>
                        <th>Valor</th>
                        <th>Nom</th>
                        <th>Min</th>
                        <th>Max</th>
                        <th>Opera</th>
                        <th>Uni</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>Potencia</td>
                        <th>{{ identifica.potencia.nominal }}</th>
                        <th>{{ identifica.potencia.minimo }}</th>
                        <th>{{ identifica.potencia.maximo }}</th>
                        <th class="{'text-primary':true}" >{{ (opera.voltaje * opera.corriente)/1000}}</th>
                        <th>{{ identifica.potencia.um }}</th>
                    </tr>
                    <tr>
                        <td>Voltaje</td>
                        <th>{{ identifica.voltaje.nominal }}</th>
                        <th>{{ identifica.voltaje.minimo }}</th>
                        <th>{{ identifica.voltaje.maximo }}</th>
                        <th :class="{'text-danger':true}">{{ opera.voltaje }}</th>
                        <th>{{ identifica.voltaje.um }}</th>
                    </tr>
                    <tr>
                        <td>Corriente</td>
                        <th>{{ identifica.corriente.nominal }}</th>
                        <th>{{ identifica.corriente.minimo }}</th>
                        <th>{{ identifica.corriente.maximo }}</th>
                        <th :class="{'text-warning':true}" >{{ opera.corriente }}</th>
                        <th>{{ identifica.corriente.um }}</th>
                    </tr>
                    <tr>
                        <td>Caudal</td>
                        <th >{{ identifica.caudal.nominal }}</th>
                        <th >{{ identifica.caudal.minimo }}</th>
                        <th >{{ identifica.caudal.maximo }}</th>
                        <th :class="{'text-success':true}" >{{ opera.caudal }}</th>
                        <th >{{ identifica.caudal.um }}</th>
                    </tr>
                </tbody>
                <tfoot>
                    <tr>
                        <th colspan="5">{{ identifica.estatus }}</th>
                    </tr>
                </tfoot>

            </table>
        </div>
        <div class="card-footer">
            <small class="text-body-secondary">{{identifica.fechaRegistro}}</small>
        </div>
    </div>
</div>
</template>
<script>
// 
export default {
    name: 'Dispositivo',
    props: {
        identifica: {
            type: Object,
            require: true
        },
        opera: {
            type: Object,
            require: true
        }
    },
    data() {
        return {
        }
    }
}
</script>
