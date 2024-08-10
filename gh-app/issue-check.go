package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"regexp"
	"strconv"
)

type IssueCheck struct {
	client openai.Client
}

func NewIssueCheck(client openai.Client) *IssueCheck {
	return &IssueCheck{client: client}
}

const (
	ResultOK   = 0
	ResultNG   = 1
	ResultSpam = 2
)

func (s *IssueCheck) CheckIssue(title, body string) int {
	ctx := context.Background()
	resp, err := s.client.CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model: "gpt-4o",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "system",
					Content: "あなたはプログラマーの有用なアシスタントです。",
				},
				{
					Role:    "user",
					Content: "次に与えられるGithubのissueを、'内容が十分'、'内容が不十分'、'スパム'のいずれかに分類してください。以下の手順に従ってください \n 1. 与えられたGithubのissueを注意深く読み、内容を理解します。\n 2. 内容が十分であるかどうかを判断します。十分な情報が含まれている場合は、'内容が十分'と分類します。\n 3. 情報が不足している場合や、質問が不明確な場合は、'内容が不十分'と分類します。 \n 4. 明らかにスパムである場合（例えば、無関係なリンクや広告、意味をなさない文章やIssueとして適切でないものなど）は、'スパム'と分類します。\n 5. 出力結果は以下のフォーマットに従ってください。出力にはXMLタグを含めないでください。 \n '内容が十分'の場合: 0 '内容が不十分'の場合:1 \n 'スパム'の場合: 2 \n 与えられるissueは以下の通りです。\n\n---" + "Title: " + title + "\n\n" + "Body: " + body,
				},
			},
		})
	if err != nil {
		return 0
	}
	log.Println(resp.Choices[0].Message.Content)
	return extractLastNumber(resp.Choices[0].Message.Content)
	// extract number from response
}

func (s *IssueCheck) ProvideFeedback(title, body string) (string, error) {
	//            {"role": "user", "content": f"```xml \n <instruction> \n  <instructions> \n あなたはプログラマーの有用なアシスタントです。次の手順に従って、与えられたGithubのissueの内容について不十分な点を考察し、指摘してください。 \n 1. Githubのissueの内容を注意深く読み、全体の文脈を理解します。 \n 2. 内容に対して、具体的な不十分な点や改善点を考えます。例えば、情報が不足している、説明が不明瞭である、または具体例が欠けているなどです。 \n 3. 指摘した点を簡潔にまとめ、200字以内で日本語で表現します。 \n 4. 出力にはXMLタグを含めず、自然な日本語の文章として仕上げます。 \n </instructions> \n <examples> \n <example>\n <input> \n 'ユーザーがログインできない問題について、詳細なエラーメッセージを提供してください。' \n </input> \n <output> \n 'エラーメッセージの具体的な内容が記載されていないため、問題の特定が難しいです。また、どのような状況でエラーが発生するのか、再現手順が不足しています。' \n </output> \n </example> \n <example> \n <input> \n '新しい機能の提案ですが、実装方法についての具体的な説明がありません。' \n </input> \n <output> \n '提案された機能の実装方法が不明確で、具体的な手順や技術的な詳細が欠けています。これにより、開発者が実装を進める際の指針が不足しています。' \n </output> \n </example> \n <example> \n <input> \n 'バグ修正のためのプルリクエストですが、テストケースが含まれていません。' \n </input> \n <output> \n 'プルリクエストにテストケースが含まれていないため、修正内容が正しく機能するかどうか確認できません。テストの有無についての説明も不足しています。' \n </output> \n 与えられたissueは以下の通りです。 \n {data}"}

	prompt := "```xml \n <instruction> \n  <instructions> \n あなたはプログラマーの有用なアシスタントです。次の手順に従って、与えられたGithubのissueの内容について不十分な点を考察し、指摘してください。 \n 1. Githubのissueの内容を注意深く読み、全体の文脈を理解します。 \n 2. 内容に対して、具体的な不十分な点や改善点を考えます。例えば、情報が不足している、説明が不明瞭である、または具体例が欠けているなどです。 \n 3. 指摘した点を簡潔にまとめ、200字以内で日本語で表現します。 \n 4. 出力にはXMLタグを含めず、自然な日本語の文章として仕上げます。 \n </instructions> \n <examples> \n <example>\n <input> \n 'ユーザーがログインできない問題について、詳細なエラーメッセージを提供してください。' \n </input> \n <output> \n 'エラーメッセージの具体的な内容が記載されていないため、問題の特定が難しいです。また、どのような状況でエラーが発生するのか、再現手順が不足しています。' \n </output> \n </example> \n <example> \n <input> \n '新しい機能の提案ですが、実装方法についての具体的な説明がありません。' \n </input> \n <output> \n '提案された機能の実装方法が不明確で、具体具体的な手順や技術的な詳細が欠けています。これにより、開発者が実装を進める際の指針が不足しています。' \n </output> \n </example> \n <example> \n <input> \n 'バグ修正のためのプルリクエストですが、テストケースが含まれていません。' \n </input> \n <output> \n 'プルリクエストにテストケースが含まれていないため、修正内容が正しく機能するかどうか確認できません。テストの有無についての説明も不足しています。' \n </output> \n 与えられたissueは以下の通りです。 \n\n"
	prompt = fmt.Sprintf("%sTitle: %s\n\nBody: %s", prompt, title, body)
	ctx := context.Background()
	resp, err := s.client.CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model: "gpt-3.5-turbo",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "system",
					Content: "あなたはプログラマーの有用なアシスタントです。",
				},
				{
					Role:    "user",
					Content: prompt,
				},
			},
		})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func extractLastNumber(input string) int {
	// 正規表現パターンを定義: 数字にマッチ
	re := regexp.MustCompile(`\d`)

	// 文字列内のすべての数字を探す
	numbers := re.FindAllString(input, -1)

	// 数字が見つからなかった場合は空文字を返す
	if len(numbers) == 0 {
		return 0
	}

	num, _ := strconv.Atoi(numbers[len(numbers)-1])
	// 最後の数字を返す
	return num
}
