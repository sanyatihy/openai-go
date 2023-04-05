package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *openAIClient) doRequest(ctx context.Context, method, endpoint string, requestData interface{}) (*http.Response, error) {
	var reqBody []byte
	var err error

	if requestData != nil {
		reqBody, err = json.Marshal(requestData)
		if err != nil {
			return nil, fmt.Errorf("error encoding request body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	c.setDefaultHeaders(req)

	return c.httpClient.Do(req)
}

func (c *openAIClient) setDefaultHeaders(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("OpenAI-Organization", c.orgID)
	req.Header.Set("Content-Type", "application/json")
}
