<template>
    <div class="analisis"
    style="flex: 1; min-width: 750px; border-right: 1px solid #ccc; padding-right: 20px; overflow-y: auto; max-height: 80vh;">

        <div class="container text-center">
            <div class="row">
                <div class="col">
                    <p class="h4">{{ $route.query.nombre }}</p>
                </div>
                <div class="col">
                    <p class="h4">Estado</p>
                    <div class="alert alert-success" role="alert">
                        {{ $route.query.estado }}
                    </div>
                </div>
            </div>
            <div class="row">
                <div class="col">
                    <p class="h5">Ubicación: {{ $route.query.ubicacion }}</p>
                </div>
            </div>
        </div>

        <div class="accordion" id="accordionPanelsStayOpenExample">
            <div class="accordion-item">
                <h2 class="accordion-header">
                    <button class="accordion-button" type="button" data-bs-toggle="collapse"
                        data-bs-target="#panelsStayOpen-collapseOne" aria-expanded="true"
                        aria-controls="panelsStayOpen-collapseOne">
                        Potencia
                    </button>
                </h2>
                <div id="panelsStayOpen-collapseOne" class="accordion-collapse collapse show">
                    <div class="accordion-body">
                        <table class="table table-hover">
                            <tr>
                                <th>Nominal</th>
                                <th>Máxima</th>
                                <th>Mínima</th>
                                <th>Unidades</th>
                            </tr>
                            <tr>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td></td>
                            </tr>
                        </table>
                    </div>
                </div>
            </div>
            <div class="accordion-item">
                <h2 class="accordion-header">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse"
                        data-bs-target="#panelsStayOpen-collapseTwo" aria-expanded="false"
                        aria-controls="panelsStayOpen-collapseTwo">
                        Voltaje
                    </button>
                </h2>
                <div id="panelsStayOpen-collapseTwo" class="accordion-collapse collapse">
                    <div class="accordion-body">
                        <table class="table table-hover">
                            <tr>
                                <th>Nominal</th>
                                <th>Máxima</th>
                                <th>Mínima</th>
                                <th>Unidades</th>
                            </tr>
                            <tr>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td></td>
                            </tr>
                        </table>
                    </div>
                </div>
            </div>
            <div class="accordion-item">
                <h2 class="accordion-header">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse"
                        data-bs-target="#panelsStayOpen-collapseThree" aria-expanded="false"
                        aria-controls="panelsStayOpen-collapseThree">
                        Corriente
                    </button>
                </h2>
                <div id="panelsStayOpen-collapseThree" class="accordion-collapse collapse">
                    <div class="accordion-body">
                        <table class="table table-hover">
                            <tr>
                                <th>Nominal</th>
                                <th>Máxima</th>
                                <th>Mínima</th>
                                <th>Unidades</th>
                            </tr>
                            <tr>
                            <tr>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td></td>
                            </tr>

                            </tr>
                        </table>
                    </div>
                </div>
            </div>
            <div class="accordion-item">
                <h2 class="accordion-header">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse"
                        data-bs-target="#panelsStayOpen-collapseFour" aria-expanded="false"
                        aria-controls="panelsStayOpen-collapseThree">
                        Caudal
                    </button>
                </h2>
                <div id="panelsStayOpen-collapseFour" class="accordion-collapse collapse">
                    <div class="accordion-body">
                        <table class="table table-hover">
                            <tr>
                                <th>Nominal</th>
                                <th>Máxima</th>
                                <th>Mínima</th>
                                <th>Unidades</th>
                            </tr>
                            <tr>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td></td>
                            </tr>

                        </table>
                    </div>
                </div>
            </div>
        </div>



    </div>

</template>
