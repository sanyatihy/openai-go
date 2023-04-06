package openai

import "fmt"

type APIError struct {
	StatusCode int
	Type       string
	Message    string
	Param      string
	Code       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error, status code: %d, type: %s, message: %s, param: %s, code: %s",
		e.StatusCode, e.Type, e.Message, e.Param, e.Code)
}

type InternalError struct {
	Message string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("Internal Error, message: %s", e.Message)
}
