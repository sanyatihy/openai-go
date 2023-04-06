package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sanyatihy/openai-go/pkg/openai"
	"go.uber.org/zap"
)

type App struct {
	logger *zap.Logger
}

func newApp() *App {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}

	return &App{
		logger: logger,
	}
}

func main() {
	app := newApp()
	defer app.logger.Sync()

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

	response, err := client.GetModel(ctx, "gpt-3.5-turbo")
	if err != nil {
		app.logger.Error("Error", zap.Error(err))
	}

	app.logger.Info(fmt.Sprintf("Got response: %s", response.ID))
}
