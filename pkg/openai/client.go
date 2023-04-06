package openai

const baseURL = "https://api.openai.com/v1"

type openAIClient struct {
	httpClient httpClient
	apiKey     string
	orgID      string
}

func NewClient(httpClient httpClient, apiKey, orgID string) Client {
	return &openAIClient{
		httpClient: httpClient,
		apiKey:     apiKey,
		orgID:      orgID,
	}
}
