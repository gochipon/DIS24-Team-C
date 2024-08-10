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
				Role:    "system",
				Content: "You are a helpful assistant of programmer.",
			},
			{
				Role:    "user",
				Content: "Summarize the following text: \n\n" + target,
			},
		},
	})
	if err != nil {
		return "", err
	}
	return comp.Choices[0].Message.Content, nil
}
