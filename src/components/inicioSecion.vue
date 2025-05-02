<template>
  <div class="login-container">
    <div class="logo-section">
      <img src="../assets/logoGatoAgua.png" alt="Logo" class="logo-image">
    </div>
    <div class="form-section">
      <div class="login-card">
        <h2 class="text-center mb-4">Iniciar Sesión</h2>
        <!-- Mensaje de alerta -->
        <div v-if="alerta.mensaje" class="alert" :class="alerta.tipo">
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
        this.$router.push('/principal');
      } else {
        this.mostrarAlerta('Usuario o contraseña incorrectos', 'alert-danger');
        this.limpiarControles();
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

.alert {
  padding: 0.75rem 1.25rem;
  margin-bottom: 1rem;
  border: 1px solid transparent;
  border-radius: 0.25rem;
}

.alert-success {
  color: #155724;
  background-color: #d4edda;
  border-color: #c3e6cb;
}

.alert-danger {
  color: #721c24;
  background-color: #f8d7da;
  border-color: #f5c6cb;
}

@media (max-width: 768px) {
  .login-container {
    flex-direction: column;
  }

  .logo-section,
  .form-section {
    width: 100%;
    min-height: 50vh;
  }
}
</style>