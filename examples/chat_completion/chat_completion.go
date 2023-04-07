package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sanyatihy/openai-go/pkg/openai"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	orgID := os.Getenv("OPENAI_ORG_ID")

	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	client := openai.NewClient(httpClient, apiKey, orgID)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	model := "gpt-3.5-turbo"
	content := "Who won the world series in 2020?"

	response, err := client.ChatCompletion(ctx, &openai.ChatCompletionRequest{
		Model: model,
		Messages: []openai.Message{
			{
				Role:    "user",
				Content: content,
			},
		},
		N:         1,
		Stream:    false,
		MaxTokens: 1024,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response.Choices[0].Message.Content)
}
