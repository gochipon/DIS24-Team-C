import json

import modules.SearchFromVDB as vectorsearch
import modules.SearchFromProgreSQL as itemsearch
import modules.Summarizer as summarizer

# set query
query = "fix a bug"
repository = "cli/cli"
topk = 3
index_name = "cli-cli-1"

# vector search
vecsearch_results = vectorsearch.main(query, topk, index_name, repository)
# sql search
sql_query = itemsearch.main(vecsearch_results)
# extract body from sql query
body_list = [x['body'] for x in sql_query]
# merge saerch results
merge_result = "\n".join(body_list)
# summarize
summary = summarizer.main(merge_result)
print(summary)