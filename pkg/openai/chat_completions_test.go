package openai

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestChatCompletions(t *testing.T) {

	tests := []struct {
		name           string
		requestOptions *ChatCompletionsRequest
		mockResponse   *http.Response
		mockError      error
		expectedResult *ChatCompletionsResponse
		expectedError  error
	}{
		{
			name: "Success",
			requestOptions: &ChatCompletionsRequest{
				Model: "some-model",
				Messages: []Message{
					{
						Role:    "system",
						Content: "You are a helpful assistant.",
					},
					{
						Role:    "user",
						Content: "Who won the world series in 2020?",
					},
				},
			},
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(bytes.NewReader([]byte(`{
					"id": "some-model",
					"object": "chat.completion",
					"created": 1677652288,
					"choices": [{
						"index": 0,
						"message": {
							"role": "assistant",
							"content": "The Los Angeles Dodgers won the World Series in 2020."
						},
						"finish_reason": "stop"
					}],
					"usage": {
						"prompt_tokens": 56,
						"completion_tokens": 31,
						"total_tokens": 87
					}
				}`))),
			},
			mockError: nil,
			expectedResult: &ChatCompletionsResponse{
				ID:      "some-model",
				Object:  "chat.completion",
				Created: 1677652288,
				Choices: []Choice{
					{
						Index: 0,
						Message: Message{
							Role:    "assistant",
							Content: "The Los Angeles Dodgers won the World Series in 2020.",
						},
						FinishReason: "stop",
					},
				},
				Usage: Usage{
					PromptTokens:     56,
					CompletionTokens: 31,
					TotalTokens:      87,
				},
			},
			expectedError: nil,
		},
		{
			name: "Error",
			requestOptions: &ChatCompletionsRequest{
				Model: "some-model",
				Messages: []Message{
					{
						Role:    "system",
						Content: "You are a helpful assistant.",
					},
					{
						Role:    "user",
						Content: "Who won the world series in 2020?",
					},
				},
			},
			mockResponse:   nil,
			mockError:      errors.New("err"),
			expectedResult: nil,
			expectedError: &InternalError{
				Message: fmt.Sprintf("error making request: %s", errors.New("err")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTPClient := new(MockHTTPClient)

			mockHTTPClient.On("Do", mock.Anything).Return(tt.mockResponse, tt.mockError)

			mockClient := NewClient(mockHTTPClient, "", "")

			response, err := mockClient.ChatCompletions(context.Background(), tt.requestOptions)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, response)
			}

			mockHTTPClient.AssertCalled(t, "Do", mock.Anything)
		})
	}
}
