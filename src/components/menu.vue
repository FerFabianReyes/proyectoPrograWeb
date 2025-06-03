<template>
    <div class="body">
        <nav class="navbar navbar-expand-lg bg-body-tertiary">
            <div class="container-fluid">
                <a class="navbar-brand" href="/menu">Gatagua</a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                    data-bs-target="#navbarNavAltMarkup" aria-controls="navbarNavAltMarkup" aria-expanded="false"
                    aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarNavAltMarkup">
                    <div class="navbar-nav">
                        <router-link to="/menu/agregarBomba" class="nav-link">Agregar Bomba</router-link>
                        <a class="nav-link" @click="cerrarSecion">Cerrar Sesión</a>
                    </div>
                </div>
            </div>
        </nav>

        <router-view></router-view>

    </div>

</template>

<script>
export default {
    name: 'Dispositivo',

    methods: {
        cerrarSecion() {
            localStorage.removeItem('usuarioActivo');
            this.$router.push('/');
        }
    }
}
</script>

<style scoped>
.navbar {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    z-index: 1000;
}

.body {
    padding-top: 56px;
}
</style>