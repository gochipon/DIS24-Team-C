<template>
  <div class="search-box">
    <Textarea v-model="searchQuery" rows="5" placeholder="Search issues and pull requests..."/>
    <Button label="Search" icon="pi pi-search" class="search-button"/>
    <div class="result-list">
      <div v-for="(item, index) in props.items as SearchResponse[]" :key="index" class="result-item">
        <template v-if="item.type === 'issue'">
          <div class="item-header">
            <span :class="['item-type', item.type]">issue</span>
<!--            open in new tab-->
            <a :href="(item.content as IssueFullResponse).issue.html_url"
               class="item-title">
              {{ (item.content as IssueFullResponse).issue.title }}
            </a>
          </div>
          <div class="item-summary">
            {{ item.summary }}
          </div>
          <div class="item-meta">
              #{{ (item.content as IssueFullResponse).issue.number }}
              <i icon="pi pi-user" :value="getUsername((item.content as IssueFullResponse).issue.user.avatar_url)"/>
              opened by {{ getUsername((item.content as IssueFullResponse).issue.user.html_url) }} &bull; {{ (item.content as IssueFullResponse).issue.comments }} comments
          </div>
        </template>
        <template v-else-if="item.type === 'pull_request'">
          <div class="item-header">
            <span :class="['item-type', item.type]">pull request</span>
            <a :href="(item.content as PullRequestResponse).html_url"
               class="item-title">
              {{ (item.content as PullRequestResponse).title }}
            </a>
          </div>
          <div class="item-summary">
            {{ item.summary }}
          </div>
          <div class="item-meta">
            #{{ (item.content as PullRequestResponse).number }}
            opened by {{ getUsername((item.content as PullRequestResponse).author) }} &bull; {{ (item.content as PullRequestResponse).comments }} comments
          </div>
        </template>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type {IssueFullResponse, IssueResponse} from "~/model/model";
import type {SearchResponse} from "~/model/search";
import type {PullRequestResponse} from "~/model/pull";

const props = defineProps({
  items: {
    type: Array<SearchResponse>,
    required: true
  }
})


const emits = defineEmits(
    [
      'update:query'
    ]
)

const searchQuery = ref('')
watch(searchQuery, () => {
  emits('update:query', searchQuery.value)
})

const isIssueResponse = (content: IssueFullResponse | PullRequestResponse): content is IssueResponse => {
  console.log(content)
  return (content as IssueFullResponse).issue.number !== undefined;
}

const getUsername = (htmlUrl: string): string => {
  return htmlUrl?.split("/")[3];
}
</script>

<style scoped>
.search-box {
  display: flex;
  flex-direction: column;
  gap: 10px;
  max-width: 800px;
  margin: 0 auto;
}

.search-button {
  align-self: flex-end;
}

.result-list {
  margin-top: 20px;
  border-top: 1px solid #e1e4e8;
}

.result-item {
  padding: 15px 0;
  border-bottom: 1px solid #e1e4e8;
}

.item-header {
  display: flex;
  align-items: center;
  margin-bottom: 5px;
}

.item-type {
  padding: 2px 7px;
  border-radius: 2em;
  font-size: 0.75rem;
  font-weight: 600;
  margin-right: 10px;
  display: inline-block;
}

.item-type.Issue {
  background-color: #d73a49;
  color: white;
}

.item-type.Pull\ Request {
  background-color: #28a745;
  color: white;
}

.item-title {
  font-size: 1rem;
  color: #0366d6;
  text-decoration: none;
}

.item-title:hover {
  text-decoration: underline;
}

.item-meta {
  font-size: 0.875rem;
  color: #586069;
}

.result-item:hover {
  background-color: #f6f8fa;
}

textarea{
  border: black 1px solid;
  border-radius: 5px;
  padding: 5px;
}
</style>
