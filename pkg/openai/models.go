package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func (c *openAIClient) GetModel(ctx context.Context, modelID string) (*ModelResponse, error) {
	url := fmt.Sprintf("%s/models/%s", baseURL, modelID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		c.logger.Error("Error making request", zap.Error(err))
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("Error reading response body", zap.Error(err))
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var response ModelResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		c.logger.Error("Error decoding response body", zap.Error(err))
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &response, nil
}
