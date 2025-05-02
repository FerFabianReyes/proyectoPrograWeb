<template>
    <div class="login-container">
      <div class="logo-section">
        <img src="../assets/logoGatoAgua.png" alt="Logo" class="logo-image">
      </div>
      <div class="form-section">
        <div class="login-card">
          <h2 class="text-center mb-4">Registro</h2>
          <form @submit.prevent="registrar">
            <div class="mb-3">
              <label for="nombre" class="form-label">Nombre</label>
              <input type="text" class="form-control" id="nombre" v-model="nombre" required>
            </div>
            <div class="mb-3">
              <label for="email" class="form-label">Correo Electrónico</label>
              <input type="email" class="form-control" id="email" v-model="email" required>
            </div>
            <div class="mb-3">
              <label for="password" class="form-label">Contraseña</label>
              <input type="password" class="form-control" id="password" v-model="password" required>
            </div>
            <div class="mb-3">
              <label for="confirmPassword" class="form-label">Confirmar Contraseña</label>
              <input type="password" class="form-control" id="confirmPassword" v-model="confirmPassword" required>
            </div>
            <button type="submit" class="btn btn-primary w-100">Registrarse</button>
          </form>
          <p class="mt-3 text-center">
            ¿Ya tienes una cuenta? <router-link to="/inicioSecion">Inicia Sesión</router-link>
          </p>
        </div>
      </div>
    </div>
  </template>
  
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