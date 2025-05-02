<template>
    <form class="row g-3 needs-validation" novalidate>
        <p class="h2 text-secondary">Regístrate</p>
        <br>
        <div class="col-md-4 position-relative">
            <label for="validationTooltip01" class="form-label">Nombre</label>
            <input type="text" class="form-control" id="validationTooltip01" value="" required>
        </div>

        <div class="col-md-4 position-relative">
            <label for="validationTooltipEmail" class="form-label">Correo</label>
            <div class="input-group has-validation">
                <input type="email" class="form-control" id="validationTooltipUsername"
                    aria-describedby="validationTooltipUsernamePrepend" required>
            </div>
        </div>
        <div class="col-md-6 position-relative">
            <label for="validationTooltip03" class="form-label">Contraseña</label>
            <input type="text" class="form-control" id="validationTooltip03" required>
            <br>
            <div class="btn-group custom-btn-group" role="group" aria-label="Basic outlined example">
                <button type="button" class="btn btn-outline-primary">Iniciar Seción</button>
                <button type="button" class="btn btn-outline-primary">Registrarse</button>
            </div>
        </div>

    </form>
</template>
