<template>
  <div class="login-container">
    <div class="logo-section">
      <img src="../assets/logoGatoAgua.png" alt="Logo" class="logo-image">
    </div>
    <div class="form-section">
      <div class="login-card">
        <h2 class="text-center mb-4">Iniciar Sesión</h2>
        <!-- Mensaje de alerta con clases de Bootstrap -->
        <div v-if="alerta.mensaje" class="alert" :class="alerta.tipo" role="alert">
          {{ alerta.mensaje }}
        </div>
        <form @submit.prevent="validar">
          <div class="mb-3">
            <label for="email" class="form-label">Correo Electrónico</label>
            <input type="email" class="form-control" id="email" v-model="login.email" required>
          </div>
          <div class="mb-3">
            <label for="password" class="form-label">Contraseña</label>
            <input type="password" class="form-control" id="password" v-model="login.password" required>
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
        email: '',
        password: ''
      },
      alerta: {
        mensaje: '',
        tipo: ''
      }
    };
  },
  methods: {
    validar() {
      const usuarios = JSON.parse(localStorage.getItem('usuarios')) || [];

      const encontrado = usuarios.find(
        u => u.email === this.login.email && u.password === this.login.password
      );

      if (encontrado) {
        localStorage.setItem('usuarioActivo', JSON.stringify(encontrado));
        this.$router.push('/menu');
      } else {
        this.mostrarAlerta('Usuario o contraseña incorrectos', 'alert-danger');
        this.login.password = ''; // Limpiar solo la contraseña por seguridad
      }
    },
    mostrarAlerta(mensaje, tipo) {
      this.alerta.mensaje = mensaje;
      this.alerta.tipo = tipo;

      // Limpiar alerta después de 5 segundos
      setTimeout(() => {
        this.alerta.mensaje = '';
        this.alerta.tipo = '';
      }, 5000);
    }
  }
};
</script>

<style scoped>

.login-container {
  display: flex;
  width: 100%;
  min-height: 100vh;
}

.logo-section {
  width: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #ffffff;
}

.form-section {
  width: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #ffffff;
}

.logo-image {
  max-width: 80%;
  max-height: 80vh;
  object-fit: contain;
  padding: 20px;
}

.login-card {
  width: 100%;
  max-width: 400px;
  padding: 2rem;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

@media (max-width: 768px) {
  .login-container {
    flex-direction: column;
  }
  
  .logo-section, .form-section {
    width: 100%;
    min-height: 50vh;
  }
}
</style>