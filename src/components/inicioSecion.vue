<template>
  <div class="login-container">
    <div class="logo-section">
      <img src="../assets/logoGatoAgua.png" alt="Logo" class="logo-image">
    </div>
    <div class="form-section">
      <div class="login-card">
        <h2 class="text-center mb-4">Iniciar Sesión</h2>
        <form @submit.prevent="iniciarSesion">
          <div class="mb-3">
            <label for="email" class="form-label">Correo Electrónico</label>
            <input type="email" class="form-control" id="email" v-model="email" required>
          </div>
          <div class="mb-3">
            <label for="password" class="form-label">Contraseña</label>
            <input type="password" class="form-control" id="password" v-model="password" required>
          </div>
          <button type="submit" class="btn btn-primary w-100">Iniciar Sesión</button>
        </form>
        <p class="mt-3 text-center">
          ¿No tienes una cuenta? <router-link to="/registro">Regístrate</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'Login',
    data() {
      return {
        login: {
          userName: '',
          password: ''
        },
        alerta: {
          mensaje: '',
        }
      };
    },
    methods: {
      validar() {
        const usuarios = JSON.parse(localStorage.getItem('usuarios')) || [];

        const encontrado = usuarios.find(
          u => u.usuario === this.login.userName && u.contrasena === this.login.password
        );

        if (encontrado) {
          this.$router.push('/Menu'); // Redirigir a la página de menú
        } else {
          this.alerta.mensaje = 'Usuario o contraseña incorrectos';
          this.limpiarControles();
        }
      }
    }
  };
</script>

<style scoped>
.login-container {
  display: flex;
  width: 100%;
}

.logo-section {
  width: 50%;
  display: flex;
  align-items: center;
}

.form-section {
  width: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #ffffff;
}

.logo-image {
  width: 100%;
  height: 100%;
  object-fit: contain;
  padding: 50px;
}

.login-card {
  width: 100%;
  padding: 30px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}
</style>