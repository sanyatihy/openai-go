package openai

import (
	"context"
	"fmt"
	"net/http"
)

func (c *openAIClient) ChatCompletion(ctx context.Context, requestOptions *ChatCompletionRequest) (*ChatCompletionResponse, error) {
	url := fmt.Sprintf("%s/chat/completions", baseURL)

	resp, err := c.doRequest(ctx, http.MethodPost, url, requestOptions)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := c.checkStatusCode(resp); err != nil {
		return nil, err
	}

	var response ChatCompletionResponse
	if err := c.processResponseBody(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
