package openai

type ChatCompletionsRequest struct {
	Model            string      `json:"model"`
	Messages         []Message   `json:"messages"`
	Temperature      float64     `json:"temperature,omitempty"`
	TopP             float64     `json:"top_p,omitempty"`
	N                int         `json:"n,omitempty"`
	Stream           bool        `json:"stream,omitempty"`
	Stop             interface{} `json:"stop,omitempty"`
	MaxTokens        int         `json:"max_tokens,omitempty"`
	PresencePenalty  float64     `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64     `json:"frequency_penalty,omitempty"`
	LogitBias        interface{} `json:"logit_bias,omitempty"`
	User             string      `json:"user,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionsResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ModelResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	OwnedBy string `json:"owned_by"`
	Root    string `json:"root"`
}
