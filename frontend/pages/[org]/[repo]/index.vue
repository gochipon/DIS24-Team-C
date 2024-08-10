<template>
  <div class="dashboard">
    <Sidebar/>
    <div class="content">
      <TopBar/>
      <div class="main-content">
        <SearchBox :items="items" @update:query="updateQuery"/>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type {IssueResponse} from "~/model/model";
import type {PullRequestResponse} from "~/model/pull";
import type {SearchResponse} from "~/model/search";
// import {debounce} from "lodash";

const route = useRoute()
const org = route.params.org
const repo = route.params.repo

import Textarea from 'primevue/textarea'
import Button from 'primevue/button'

const items = ref<SearchResponse[]>([])


const searchQuery = ref('')


// get current time
let lastrun = Date.now()
let running = false
const updateQuery = async (newQuery: string) => {
  const now = Date.now()
  if (now - lastrun < 500) {
    return
  }
  lastrun = now
  if (running || newQuery === "") {
    return
  }
  running = true
  console.log("running")
  const {data, error} = await useFetch<SearchResponse[]>(
      'https://api.github-tracker.dev' + `/api/v1/${org}/${repo}/search`,
      {
        method: 'POST',
        body: {query: newQuery}
      }
  )
  running = false
  if (error.value) {
    console.log(error.value)
    return
  }
  items.value = data.value
}

</script>

<style scoped>
.dashboard {
  display: flex;
  height: 100vh;
}

.content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.main-content {
  padding: 20px;
  flex: 1;
  overflow-y: auto;
}
</style>
