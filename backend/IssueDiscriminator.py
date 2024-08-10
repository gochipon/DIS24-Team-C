import openai
import json

config = json.load(open('config.json'))
OPEN_AI_API_KEY = config["OPEN_AI_API_KEY"]
openai.api_key = OPEN_AI_API_KEY


def discriminate(data: str) -> str:
    completion = openai.chat.completions.create(
        model="gpt-3.5-turbo", 
        messages=[
            {"role": "system", "content": "あなたはプログラマーの有用なアシスタントです。"},
            {"role": "user", "content": f"次に与えられるGithubのissueを、'内容が十分'、'内容が不十分'、'スパム'のいずれかに分類してください。以下の手順に従ってください \n 1. 与えられたGithubのissueを注意深く読み、内容を理解します。\n 2. 内容が十分であるかどうかを判断します。十分な情報が含まれている場合は、'内容が十分'と分類します。\n 3. 情報が不足している場合や、質問が不明確な場合は、'内容が不十分'と分類します。 \n 4. 明らかにスパムである場合（例えば、無関係なリンクや広告、意味をなさない文章など）は、'スパム'と分類します。\n 5. 出力結果は以下のフォーマットに従ってください。出力にはXMLタグを含めないでください。 \n '内容が十分'の場合: 0 '内容が不十分'の場合:1 \n 'スパム'の場合: 2 \n 与えられるissueは以下の通りです。 \n {data}"}
        ])
    
    result = completion.choices[0].message.content
    return result

def main(data: str) -> str:
    return discriminate(data)

