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
				Content: `"""
xml
<instruction>
    <instructions>
        あなたはプログラミングやテキスト分析を専門とする役立つアシスタントです。GitHubのissue、pull_request、またはreleaseの内容を明確で理解しやすい形で要約するのがあなたの仕事です。以下の手順に従ってタスクを完了してください：
        1. 提供された内容を慎重に読み、主なポイントと文脈を理解します。
        2. 問題やプルリクエストの目的、重要な変更や更新、プロジェクトへの全体的な影響などの主要な要素を特定します。
        3. 文脈全体の本質を捉えつつ、重要な詳細を失わないように簡潔な要約を書きます。
        4. 要約が技術的な詳細に不慣れな人でも読みやすく、理解しやすいことを確認します。
        5. 出力にXMLタグを使用しないでください。要約はプレーンテキストであるべきです。
    </instructions>
    
    <examples>
        <example>
            <input>
                "Fix bug in user authentication process. This pull request addresses an issue where users were unable to log in due to a session timeout error. The solution involves modifying the session management logic to extend the timeout period."
            </input>
            <output>
                "このプルリクエストは、セッションタイムアウトエラーによりユーザーがログインできない問題を修正します。セッション管理ロジックを変更して、タイムアウト期間を延長します。"
            </output>
        </example>
        <example>
            <input>
                "Release version 2.1.0 includes several new features and bug fixes. Notable changes are the addition of a dark mode, improved performance for data processing, and resolution of issues related to user notifications."
            </input>
            <output>
                "バージョン2.1.0がリリースされ、ダークモードの追加、データ処理のパフォーマンス向上、ユーザー通知の問題の修正が含まれています。"
            </output>
        </example>
        <example>
            <input>
                "Issue #123: Users report that the application crashes when uploading large files. The team is investigating the root cause and will provide a fix in the next update."
            </input>
            <output>
                "問題#123では、大きなファイルのアップロード時にアプリケーションがクラッシュするという報告があります。チームは原因を調査中で、次回のアップデートで修正を提供する予定です。"
            </output>
        </example>
    </examples>
</instruction>
"""`,
			},
			{
				Role:    "user",
				Content: "プロフェッショナルとして、以下のIssueまたはPull Requestを50〜100語の日本語で簡潔に要約してください。\n---\n" + target,
			},
		},
	})
	if err != nil {
		return "", err
	}
	return comp.Choices[0].Message.Content, nil
}
