package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *openAIClient) doRequest(ctx context.Context, method, endpoint string, requestData interface{}) (*http.Response, error) {
	var reqBody bytes.Buffer

	if requestData != nil {
		encoder := json.NewEncoder(&reqBody)
		if err := encoder.Encode(requestData); err != nil {
			return nil, &InternalError{
				Message: fmt.Sprintf("error encoding request body: %s", err),
			}
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, &reqBody)
	if err != nil {
		return nil, &InternalError{
			Message: fmt.Sprintf("error creating request: %s", err),
		}
	}

	c.setDefaultHeaders(req)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &InternalError{
			Message: fmt.Sprintf("error making request: %s", err),
		}
	}

	return res, nil
}

func (c *openAIClient) setDefaultHeaders(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("OpenAI-Organization", c.orgID)
	req.Header.Set("Content-Type", "application/json")
}

func (c *openAIClient) processResponseBody(resp *http.Response, target interface{}) error {
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(target); err != nil {
		return &InternalError{
			Message: fmt.Sprintf("error decoding response body: %s", err),
		}
	}

	return nil
}

func (c *openAIClient) checkStatusCode(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusTooManyRequests:
		return c.extractAPIError(resp)
	case http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound:
		return c.extractAPIError(resp)
	case http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		return c.extractAPIError(resp)
	default:
		return c.extractAPIError(resp)
	}
}

func (c *openAIClient) extractAPIError(resp *http.Response) error {
	var apiErrorBody struct {
		Error APIError `json:"error"`
	}

	if err := c.processResponseBody(resp, &apiErrorBody); err != nil {
		return err
	}

	apiErrorBody.Error.StatusCode = resp.StatusCode
	return &apiErrorBody.Error
}
