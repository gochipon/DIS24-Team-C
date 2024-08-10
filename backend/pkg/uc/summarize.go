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
				Content: `
<instruction>
    <instructions>
        You are a helpful assistant specializing in programming and text analysis. Your task is to summarize the content of issues, pull requests, or releases from GitHub in a clear and understandable manner. Follow these steps to complete the task:
        1. Read the provided content carefully to understand the main points and context.
        2. Identify key elements such as the purpose of the issue or pull request, any significant changes or updates, and the overall impact on the project.
        3. Write a concise summary that captures the essence of the content without losing important details.
        4. Ensure that the summary is easy to read and understand for someone who may not be familiar with the specific technical details.
        5. Avoid using any XML tags in your output; the summary should be plain text.
    </instructions>
    
    <examples>
        <example>
            <input>
                "Fix bug in user authentication process. This pull request addresses an issue where users were unable to log in due to a session timeout error. The solution involves modifying the session management logic to extend the timeout period."
            </input>
            <output>
                "This pull request fixes a bug in the user authentication process that prevented users from logging in due to session timeout errors. It modifies the session management logic to extend the timeout period."
            </output>
        </example>
        <example>
            <input>
                "Release version 2.1.0 includes several new features and bug fixes. Notable changes are the addition of a dark mode, improved performance for data processing, and resolution of issues related to user notifications."
            </input>
            <output>
                "Version 2.1.0 has been released, featuring new additions such as dark mode, enhanced performance for data processing, and fixes for user notification issues."
            </output>
        </example>
        <example>
            <input>
                "Issue #123: Users report that the application crashes when uploading large files. The team is investigating the root cause and will provide a fix in the next update."
            </input>
            <output>
                "Issue #123 reports that the application crashes during large file uploads. The team is currently investigating the cause and plans to provide a fix in the next update."
            </output>
        </example>
    </examples>
</instruction>
`,
			},
			{
				Role:    "user",
				Content: "Please act as a professional and summarize the " + target + " in a concise manner of 50-100 words.",
			},
		},
	})
	if err != nil {
		return "", err
	}
	return comp.Choices[0].Message.Content, nil
}
