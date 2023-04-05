package openai

import (
	"context"
	"net/http"
)

type Client interface {
	ChatCompletions(ctx context.Context, requestOptions *ChatCompletionsRequest) (*ChatCompletionsResponse, error)
	GetModel(ctx context.Context, modelID string) (*ModelResponse, error)
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
