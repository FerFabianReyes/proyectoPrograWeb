<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
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
