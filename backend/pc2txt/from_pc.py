from pinecone.grpc import PineconeGRPC as Pinecone_api
from langchain_openai import OpenAIEmbeddings 
from langchain_pinecone import PineconeVectorStore  
from langchain_openai import ChatOpenAI  
from langchain.chains import RetrievalQA  
import time  
import os

def main(query: str, topk: int) -> None:  
    # configure client  
    PINECONE_API_KEY = os.environ['PINECONE_API_KEY']
    OPEN_AI_API_KEY = os.environ['OPENAI_API_KEY']

    pc = Pinecone_api(api_key=PINECONE_API_KEY)  

    index_name = 'cli-cli'
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

    llm = ChatOpenAI(  
        openai_api_key=OPEN_AI_API_KEY,  
        model_name='gpt-3.5-turbo',  
        temperature=0.0  
    )  

    qa = RetrievalQA.from_chain_type(  
        llm=llm,  
        chain_type="stuff",  
        retriever=vectorstore.as_retriever()  
    )  

    vectorstore.similarity_search_with_relevance_scores(  
        query,  # our search query  
        k=topk  # return 3 most relevant docs  
    )  

    print(qa.run(query))

if __name__ == "__main__":  
    query = "Tell me information about `gh repo view` on forked repositories."
    main(query, topk=3)