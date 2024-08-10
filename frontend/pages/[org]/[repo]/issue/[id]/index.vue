<template>
  <DataTable :value="issues" class="p-datatable-gridlines">
    <Column field="title" header="Title"></Column>
    <Column field="status" header="Status">
      <template #body="slotProps">
        <Tag :severity="slotProps.data.status === 'Open' ? 'success' : 'danger'">
          {{ slotProps.data.status }}
        </Tag>
      </template>
    </Column>
    <Column field="author" header="Author"></Column>
    <Column field="created_at" header="Created At"></Column>
    <Column header="Actions">
      <template #body="slotProps">
        <Button label="View" class="p-button-text"></Button>
      </template>
    </Column>
  </DataTable>
</template>

<script setup>
import { ref } from 'vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Tag from 'primevue/tag';
import Button from 'primevue/button';

const issues = ref([
  { id: 1, title: 'Issue 1', status: 'Open', author: 'User1', created_at: '2023-08-09' },
  { id: 2, title: 'Issue 2', status: 'Closed', author: 'User2', created_at: '2023-08-08' },
  { id: 3, title: 'Issue 3', status: 'Open', author: 'User3', created_at: '2023-08-07' },
]);
</script>
