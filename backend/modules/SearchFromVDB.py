from pinecone.grpc import PineconeGRPC as Pinecone_api
from langchain_openai import OpenAIEmbeddings 
from langchain_pinecone import PineconeVectorStore  
from langchain_openai import ChatOpenAI  
# from langchain.chains import RetrievalQA  
# import time  
import os
import json

def main(query: str, topk: int, index_name: str, repository: str) -> None:  
    # configure client  
    # load environment variables
    config = json.load(open('config.json'))

    PINECONE_API_KEY = config['PINECONE_API_KEY']
    OPEN_AI_API_KEY = config['OPEN_AI_API_KEY']

    pc = Pinecone_api(api_key=PINECONE_API_KEY)  
    index = pc.Index(index_name)  

    model_name = 'text-embedding-ada-002'  
    embeddings = OpenAIEmbeddings(  
        model=model_name,  
        openai_api_key=OPEN_AI_API_KEY 
    )  

    text_field = "text"  
    vectorstore = PineconeVectorStore(  
        index, 
        embeddings, 
        text_field  
    )

    top_k_data = vectorstore.similarity_search_with_relevance_scores(  
        query,  # our search query  
        k=topk,  # return 3 most relevant docs  
        filter={
            'repository':repository,
            # '_ab_stream':'releases'
        },
    )  

    top_k_id = []

    if top_k_data:
        for k_data in top_k_data:
            
            info={}
            meta= k_data[0].metadata
            stream = meta['_ab_stream']
            info['repository'] = meta['repository']

            info['socre'] = k_data[1]
            
            if stream == 'releases':
                info['name'] = stream
                info['id'] = meta['name']
                

            elif stream == 'comments':
                info['name'] = 'issues'
                info['id'] = meta['issue_url'].split('/')[-1]        
            elif stream == 'issues':
                info['name'] = stream
                info['id'] = meta['number']

            elif stream == 'pull_requests':
                info['name'] = stream
                info['id'] = meta['number']
            elif stream == 'review_comments':
                info['name'] = 'pull_requests'
                info['id'] = meta['pull_request_url'].split('/')[-1]
            else:
                continue
        
            top_k_id.append(info)

    # print(top_k_id)
    return top_k_id
    
if __name__ == "__main__":  
    query = "Tell me information about `gh repo view` on forked repositories."
    repository = "cli/cli"
    main(query, topk=5, index_name='cli-cli-1', repository=repository)