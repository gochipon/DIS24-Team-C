import openai
import json

config = json.load(open('config.json'))
OPEN_AI_API_KEY = config["OPEN_AI_API_KEY"]
openai.api_key = OPEN_AI_API_KEY


def summarize(query:str, data: str) -> str:
    completion = openai.chat.completions.create(
        model="gpt-3.5-turbo", 
        messages=[
            {"role": "system", "content": "You are a helpful assistant of programmer."},
            {"role": "user", "content": f"Extract the most relevant part for {query} from the following texts :{data}"}
        ])
    
    result = completion.choices[0].message.content
    return result

def main(query:str, data: str) -> str:
    return summarize(query, data)