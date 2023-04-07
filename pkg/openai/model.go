package openai

import (
	"context"
	"fmt"
	"net/http"
)

func (c *openAIClient) GetModel(ctx context.Context, modelID string) (*ModelResponse, error) {
	url := fmt.Sprintf("%s/models/%s", baseURL, modelID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := c.checkStatusCode(resp); err != nil {
		return nil, err
	}

	var response ModelResponse
	if err := c.processResponseBody(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
