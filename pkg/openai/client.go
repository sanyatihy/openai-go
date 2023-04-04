package openai

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

const baseURL = "https://api.openai.com/v1"

type openAIClient struct {
	httpClient *http.Client
	logger     *zap.Logger
	apiKey     string
	orgID      string
}

func NewClient(logger *zap.Logger, apiKey, orgID string) Client {
	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}

	return &openAIClient{
		httpClient: &http.Client{
			Transport: transport,
		},
		logger: logger,
		apiKey: apiKey,
		orgID:  orgID,
	}
}
