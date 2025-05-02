<template>
  <div class="login-container">
    <div class="logo-section">
      <img src="../assets/logoGatoAgua.png" alt="Logo" class="logo-image">
    </div>
    <div class="form-section">
      <div class="login-card">
        <h2 class="text-center mb-4">Registro</h2>
        
        <!-- Mensaje de alerta -->
        <div v-if="alerta.mensaje" class="alert" :class="alerta.tipo">
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

      // Ver los usuarios registrados
      const usuariosRegistrados = JSON.parse(localStorage.getItem('usuarios')) || [];

      // Ver si el usuario ya existe
      const existe = usuariosRegistrados.some(u => u.email === email);
      if (existe) {
        this.mostrarAlerta('El correo electrónico ya está registrado', 'alert-danger');
        return;
      }

      // Crea un nuevo usuario 
      const nuevoUsuario = { 
        nombre, 
        email, 
        password
      };

      // Agregar al nuevo usuario
      usuariosRegistrados.push(nuevoUsuario);

      // Guardar en localStorage
      localStorage.setItem('usuarios', JSON.stringify(usuariosRegistrados));

      this.registro = {
        nombre: '',
        email: '',
        password: '',
        confirmPassword: ''
      };
  
      this.$router.push('/');
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
  
  .logo-section, .form-section {
    width: 100%;
    min-height: 50vh;
  }
}
</style>