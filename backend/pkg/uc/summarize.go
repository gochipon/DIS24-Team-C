package uc

import (
	"context"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/config"
	"github.com/sashabaranov/go-openai"
)

type SummarizeUseCase struct {
	client openai.Client
}

func NewSummarizeUseCase() *SummarizeUseCase {
	client := openai.NewClient(config.C.OpenaiKey)
	return &SummarizeUseCase{
		client: *client,
	}
}

func (uc *SummarizeUseCase) Summarize(target string) (string, error) {
	ctx := context.Background()
	// gpt4o-mini
	comp, err := uc.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: "gpt-4o-mini",
		Messages: []openai.ChatCompletionMessage{
			{
				Role: "system",
				Content: `xml  
<instruction>
    <instructions>
        あなたはプログラミングやテキスト分析を専門とする役立つアシスタントです。
        GitHubのissue、pull_request、またはreleaseの内容を明確で理解しやすい形で要約するのがあなたの仕事です。以下の手順に従ってタスクを完了してください：

        1. 提供された内容を慎重に読み、主なポイントと文脈を理解します。
        2. 報告された問題やプルリクエストの目的、重要な変更や更新、プロジェクトへの全体的な影響などの主要な要素を特定します。
        3. 文脈全体の本質を捉えつつ、簡潔な要約を書きます。
        4. 要約が技術的な詳細に不慣れな人でも読みやすく、理解しやすいことを確認します。
        5. 出力にXMLタグを使用しないでください。要約はプレーンテキストであるべきです。
    </instructions>
    
    <examples>
        <example>
            <input>
                "Fix bug in user authentication process. This pull request addresses an issue where users were unable to log in due to a session timeout error. The solution involves modifying the session management logic to extend the timeout period."
            </input>
            <output>
                "セッションタイムアウトエラーによりユーザーがログインできない問題を修正したと報告。セッション管理ロジックを変更して、タイムアウト期間を延長しました。"
            </output>
        </example>
        <example>
            <input>
                "Release version 2.1.0 includes several new features and bug fixes. Notable changes are the addition of a dark mode, improved performance for data processing, and resolution of issues related to user notifications."
            </input>
            <output>
                "バージョン2.1.0がリリースされ、主なものとして、ダークモードの追加、データ処理のパフォーマンス向上、ユーザー通知の問題の修正が含まれています。"
            </output>
        </example>
        <example>
            <input>
                "Issue #123: Users report that the application crashes when uploading large files. The team is investigating the root cause and will provide a fix in the next update."
            </input>
            <output>
                "大きなファイルのアップロード時にアプリケーションがクラッシュするという報告を行っています。チームは原因を調査中で、次回のアップデートで修正を提供する予定です。"
            </output>
        </example>
    </examples>
</instruction>
`,
			},
			{
				Role:    "user",
				Content: target + "を自然な日本語50〜100語で簡潔に要約してください。要約は50〜100語の範囲内で、明確で理解しやすいものである必要があります。技術的な詳細を簡潔にまとめ、一般の読者にも理解しやすいようにしてください。",
			},
		},
	})
	if err != nil {
		return "", err
	}
	return comp.Choices[0].Message.Content, nil
}
