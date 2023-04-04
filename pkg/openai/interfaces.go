package openai

import "context"

type Client interface {
	ChatCompletions(ctx context.Context, requestOptions *ChatCompletionsRequest) (*ChatCompletionsResponse, error)
	GetModel(ctx context.Context, modelID string) (*ModelResponse, error)
}
