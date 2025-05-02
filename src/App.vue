<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <div class="logo-container">
      <p class="h1 text-primary">GatAgua</p>
      <figure class="figure">
        <img src="@/assets/logoGatoAgua.png" class="figure-img img-fluid rounded">
      </figure>
    </div>
  </header>

  <RouterView />
</template>

<style scoped>

.logo-container {
  display: flex;
  flex-direction: column;
  align-items: center; 
  gap: 10px;
}

.figure-img {
  max-width: 250px; 
}


</style>
