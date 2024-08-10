<script setup lang="ts"></script>

<template>
  <div>
    <Menubar :model="[{
      label: 'Home',
      icon: 'pi pi-fw pi-home',
      to: '/'
    }]"/>
    <slot/>
  </div>
</template>

<style scoped></style>
