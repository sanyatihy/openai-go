package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (c *openAIClient) doRequest(ctx context.Context, method, endpoint string, requestData interface{}) (*http.Response, error) {
	var reqBody []byte
	var err error

	if requestData != nil {
		reqBody, err = json.Marshal(requestData)
		if err != nil {
			return nil, errors.Wrap(err, "error encoding request body")
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, errors.Wrap(err, "error creating request")
	}

	c.setDefaultHeaders(req)

	return c.httpClient.Do(req)
}

func (c *openAIClient) setDefaultHeaders(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("OpenAI-Organization", c.orgID)
	req.Header.Set("Content-Type", "application/json")
}
