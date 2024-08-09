import json

import modules.SearchFromVDB as vectorsearch
import modules.SearchFromProgreSQL as itemsearch
import modules.Summarizer as summarizer
import modules.RelatedPartFinder as relatedpartfinder

# set query
query = "fix a bug"
repository = "cli/cli"
topk = 3
index_name = "cli-cli-1"

# vector search
vecsearch_results = vectorsearch.main(query, topk, index_name, repository)
# sql search
sql_query = itemsearch.main(vecsearch_results)
result = [{"title": _res["title"], "url": _res["url"], "summary": summarizer.main(_res['body']), "related_part": relatedpartfinder.main(query, _res["body"])} for _res in sql_query]
print(result)
