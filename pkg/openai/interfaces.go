package openai

import (
	"context"
	"net/http"
)

type Client interface {
	ChatCompletion(ctx context.Context, requestOptions *ChatCompletionRequest) (*ChatCompletionResponse, error)
	GetModel(ctx context.Context, modelID string) (*ModelResponse, error)
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
