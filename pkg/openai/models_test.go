package openai

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetModel(t *testing.T) {
	tests := []struct {
		name           string
		modelID        string
		mockResponse   *http.Response
		mockStatusCode int
		mockError      error
		expectedError  error
	}{
		{
			name:    "Success",
			modelID: "some-model",
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(bytes.NewReader([]byte(`{
					"id": "gpt-3.5-turbo",
					"object": "model",
					"created": 1677610602,
					"owned_by": "openai",
					"permission": [{
						"id": "modelperm-M56FXnG1AsIr3SXq8BYPvXJA",
						"object": "model_permission",
						"created": 1679602088,
						"allow_create_engine": false,
						"allow_sampling": true,
						"allow_logprobs": true,
						"allow_search_indices": false,
						"allow_view": true,
						"allow_fine_tuning": false,
						"organization": "*",
						"group": null,
						"is_blocking": false
					}],
					"root": "gpt-3.5-turbo",
					"parent": null
				}`))),
			},
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "Error",
			modelID:       "some-model",
			mockResponse:  nil,
			mockError:     errors.New("error making request"),
			expectedError: fmt.Errorf("error making request: %w", errors.New("error making request")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTPClient := new(MockHTTPClient)

			mockHTTPClient.On("Do", mock.Anything).Return(tt.mockResponse, tt.mockError)

			apiClient := NewClient(mockHTTPClient, zap.NewNop(), "", "")

			response, err := apiClient.GetModel(context.Background(), tt.modelID)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
			}

			mockHTTPClient.AssertCalled(t, "Do", mock.Anything)
		})
	}
}
