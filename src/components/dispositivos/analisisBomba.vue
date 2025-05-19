<template>
  <div class="container text-center">
    <div class="row">
      <div class="col">
        <p class="h4">Nombre: {{ $route.query.nombre }}</p>
      </div>
      <div class="col">
        <p class="h4">Estado: {{ $route.query.estado }}</p>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <p class="h5">Ubicación: {{ $route.query.ubicacion }}</p>
      </div>
    </div>
  </div>
</template>
