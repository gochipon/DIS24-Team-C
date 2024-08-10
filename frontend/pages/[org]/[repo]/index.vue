<template>
  <div class="dashboard">
    <div class="content">
      <TopBar/>
      <h1>{{org}}/{{repo}}のベクトル検索エンジン</h1>
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

const route = useRoute()
const org = route.params.org
const repo = route.params.repo

import Textarea from 'primevue/textarea'
import Button from 'primevue/button'

const items = ref<SearchResponse[]>([])


const searchQuery = ref('')


// get current time
let lastrun = Date.now()
const updateQuery = async (newQuery: string) => {
  const now = Date.now()
  lastrun = now
  await (async () => {
    await new Promise((resolve) => {
      setTimeout(resolve, 500)
    })
  })()
  if (now !== lastrun) {
    return
  }
  if (newQuery === "") {
    return
  }
  const {data, error} = await useFetch<SearchResponse[]>(
      'https://api.github-tracker.dev' + `/api/v1/${org}/${repo}/search`,
      {
        method: 'POST',
        body: {query: newQuery}
      }
  )
  if (error.value) {
    console.log(error.value)
    return
  }
  items.value = data.value
}

</script>

<style scoped>
h1 {
  text-align: center;
}
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
