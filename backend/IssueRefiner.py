import openai
import json
import os
import glob

config = json.load(open('config.json'))
OPEN_AI_API_KEY = config["OPEN_AI_API_KEY"]
openai.api_key = OPEN_AI_API_KEY
template_dirs = glob.glob(os.path.join(".", "issue_templates", "*.md"))

templates = []
for filename in template_dirs:
    with open(filename) as f:
        templates.append(f.read())
        
def check_issue_format(data: str, templates: list) -> str:
    completion = openai.chat.completions.create(
        model="gpt-4o-mini", 
        messages=[
            {"role": "system", "content": "あなたはプログラマーの有用なアシスタントです。"},
            {"role": "user", "content": f"```xml \n <instruction> \n  <instructions> \n あなたはプログラマーの有用なアシスタントです。次の手順に従って、与えられたGithubのissueの内容について不十分な点を考察し、指摘してください。 \n 1. Githubのissueの内容を注意深く読み、全体の文脈を理解します。 \n 2. 内容に対して、具体的な不十分な点や改善点を考えます。例えば、情報が不足している、説明が不明瞭である、または具体例が欠けているなどです。 \n 3. 指摘した点を簡潔にまとめ、200字以内で日本語で表現します。 \n 4. 出力にはXMLタグを含めず、自然な日本語の文章として仕上げます。 \n </instructions> \n <examples> \n <example>\n <input> \n 'ユーザーがログインできない問題について、詳細なエラーメッセージを提供してください。' \n </input> \n <output> \n 'エラーメッセージの具体的な内容が記載されていないため、問題の特定が難しいです。また、どのような状況でエラーが発生するのか、再現手順が不足しています。' \n </output> \n </example> \n <example> \n <input> \n '新しい機能の提案ですが、実装方法についての具体的な説明がありません。' \n </input> \n <output> \n '提案された機能の実装方法が不明確で、具体的な手順や技術的な詳細が欠けています。これにより、開発者が実装を進める際の指針が不足しています。' \n </output> \n </example> \n <example> \n <input> \n 'バグ修正のためのプルリクエストですが、テストケースが含まれていません。' \n </input> \n <output> \n 'プルリクエストにテストケースが含まれていないため、修正内容が正しく機能するかどうか確認できません。テストの有無についての説明も不足しています。' \n </output> \n 与えられたissueは以下の通りです。 \n {data}"}
        ])
    
    result = completion.choices[0].message.content
    return result

def suggest_refined_issue(data: str, templates: list) -> str:
    completion = openai.chat.completions.create(
        model="gpt-4o-mini", 
        messages=[
            {"role": "system", "content": "あなたはプログラマーの有用なアシスタントです。"},
            {"role": "user", "content": f"次の手順に従って、与えられたGithubのissueに対して適切なテンプレートを選び、そのテンプレートに基づいてissueを書き直してください。 \n 1. 与えられたGithubのissueを読み、内容を理解します。 \n 2. 提供されたテンプレートの中から、最も適切なものを選択します。 \n 3. 選択したテンプレートのフォーマットに従って、与えられたissueを再構成します。 \n 4. 出力は日本語で300字程度にまとめ、XMLタグは含めないようにしてください。 \n 5. 出力内容が明確で、読みやすいことを確認してください。 \n 与えられたissueは以下の通りです。{data} \n テンプレートは以下の中から選択してください。{templates}"}
        ])
    
    result = completion.choices[0].message.content
    return result

def main(data: str) -> str:
    advise = check_issue_format(data, templates)
    refined_issue = suggest_refined_issue(data, templates)
    result = {"advise":advise, "refined_issue":refined_issue}
    return result