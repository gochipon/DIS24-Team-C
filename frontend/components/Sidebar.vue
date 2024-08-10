<template>
  <div class="sidebar">
    <Menu :model="items" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Menu } from 'primevue/menu'

const items = ref([
  { label: 'Dashboard', icon: 'pi pi-fw pi-home' },
  { label: 'Settings', icon: 'pi pi-fw pi-cog' },
  // 他のメニューアイテムを追加
])
</script>

<style scoped>
.sidebar {
  width: 250px;
  background-color: #2c3e50;
  color: white;
  padding: 20px;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  height: 100%;
}
</style>
