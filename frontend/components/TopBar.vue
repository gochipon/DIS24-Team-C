<template>
  <div class="topbar">
    <span class="title">Dashboard</span>
    <Button label="Logout" icon="pi pi-sign-out" />
  </div>
</template>

<script setup>
import Button from 'primevue/button'
</script>

<style scoped>
.topbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #34495e;
  color: white;
  padding: 10px 20px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}
</style>
