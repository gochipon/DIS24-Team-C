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

<style >
body{
  margin: 0;
  padding: 0;
}
</style>
