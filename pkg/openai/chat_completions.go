package openai

import (
	"context"
	"fmt"
	"net/http"
)

func (c *openAIClient) ChatCompletions(ctx context.Context, requestOptions *ChatCompletionsRequest) (*ChatCompletionsResponse, error) {
	url := fmt.Sprintf("%s/chat/completions", baseURL)

	resp, err := c.doRequest(ctx, http.MethodPost, url, requestOptions)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if err := c.checkStatusCode(resp); err != nil {
		return nil, err
	}

	var response ChatCompletionsResponse
	if err := c.processResponseBody(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
