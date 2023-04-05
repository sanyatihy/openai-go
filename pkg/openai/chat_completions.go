package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func (c *openAIClient) ChatCompletions(ctx context.Context, requestOptions *ChatCompletionsRequest) (*ChatCompletionsResponse, error) {
	url := fmt.Sprintf("%s/chat/completions", baseURL)

	resp, err := c.doRequest(ctx, http.MethodPost, url, requestOptions)
	if err != nil {
		c.logger.Error("Error making request", zap.Error(err))
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		c.logger.Error("Too many requests", zap.Error(err))
		return nil, fmt.Errorf("too many requests: %w", err)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("Error reading response body", zap.Error(err))
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var response ChatCompletionsResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		c.logger.Error("Error decoding response body", zap.Error(err))
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &response, nil
}
