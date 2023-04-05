package openai

import (
	"go.uber.org/zap"
)

const baseURL = "https://api.openai.com/v1"

type openAIClient struct {
	httpClient httpClient
	logger     *zap.Logger
	apiKey     string
	orgID      string
}

func NewClient(httpClient httpClient, logger *zap.Logger, apiKey, orgID string) Client {
	return &openAIClient{
		httpClient: httpClient,
		logger:     logger,
		apiKey:     apiKey,
		orgID:      orgID,
	}
}
