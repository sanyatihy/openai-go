package main

import (
	"context"
	"fmt"
	"log"
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

	client := openai.NewClient(app.logger, apiKey, orgID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.GetModel(ctx, "gpt-3.5-turbo")
	if err != nil {
		app.logger.Error("Error", zap.Error(err))
	}

	fmt.Println(response.ID)
}
