<template>
  <div class="login-container">
    <div class="logo-section">
      <img src="../assets/logoGatoAgua.png" alt="Logo" class="logo-image">
    </div>
    <div class="form-section">
      <div class="login-card">
        <h2 class="text-center mb-4">Registro</h2>
        
        <!-- Mensaje de alerta usando clases de Bootstrap -->
        <div v-if="alerta.mensaje" class="alert" :class="alerta.tipo" role="alert">
          {{ alerta.mensaje }}
        </div>
        
        <form @submit.prevent="registrar">
          <div class="mb-3">
            <label for="nombre" class="form-label">Nombre</label>
            <input type="text" class="form-control" id="nombre" v-model="registro.nombre" required>
          </div>
          <div class="mb-3">
            <label for="email" class="form-label">Correo Electrónico</label>
            <input type="email" class="form-control" id="email" v-model="registro.email" required>
          </div>
          <div class="mb-3">
            <label for="password" class="form-label">Contraseña</label>
            <input type="password" class="form-control" id="password" v-model="registro.password" required>
          </div>
          <div class="mb-3">
            <label for="confirmPassword" class="form-label">Confirmar Contraseña</label>
            <input type="password" class="form-control" id="confirmPassword" v-model="registro.confirmPassword" required>
          </div>
          <button type="submit" class="btn btn-primary w-100">Registrarse</button>
        </form>
        <p class="mt-3 text-center">
          ¿Ya tienes una cuenta? <router-link to="/">Inicia Sesión</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Registro',
  data() {
    return {
      registro: {
        nombre: '',
        email: '',
        password: '',
        confirmPassword: ''
      },
      alerta: {
        mensaje: '',
        tipo: ''
      }
    };
  },
  methods: {
    registrar() {
      const { nombre, email, password, confirmPassword } = this.registro;

      if (password !== confirmPassword) {
        this.mostrarAlerta('Las contraseñas no coinciden', 'alert-danger');
        return;
      }

      const usuariosRegistrados = JSON.parse(localStorage.getItem('usuarios')) || [];
      const existe = usuariosRegistrados.some(u => u.email === email);
      
      if (existe) {
        this.mostrarAlerta('El correo electrónico ya está registrado', 'alert-danger');
        return;
      }

      const nuevoUsuario = { nombre, email, password };
      usuariosRegistrados.push(nuevoUsuario);
      localStorage.setItem('usuarios', JSON.stringify(usuariosRegistrados));

      this.mostrarAlerta('¡Registro exitoso! Redirigiendo...', 'alert-success');
      
      this.registro = { nombre: '', email: '', password: '', confirmPassword: '' };

      setTimeout(() => {
        this.$router.push('/');
      }, 2000);
    },
    mostrarAlerta(mensaje, tipo) {
      this.alerta.mensaje = mensaje;
      this.alerta.tipo = tipo;
      
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