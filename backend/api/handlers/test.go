package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
)

type TestHandlerInput struct{}

type TestHandlerOutput struct {
	Body struct {
		Status string    `json:"status"`
		Time   time.Time `json:"time"`
	}
}

var OpTestHandler = huma.Operation{
	OperationID: "test",
	Method:      http.MethodGet,
	Path:        "/test",
	Summary:     "Test endpoint",
	Tags:        []string{"test"},
}

func TestHandler(ctx context.Context, input *TestHandlerInput) (*TestHandlerOutput, error) {
	response := &TestHandlerOutput{}
	response.Body.Status = "ok"
	response.Body.Time = time.Now()
	// app.GetApp().EntClient.Model.Query().All(ctx)
	return response, nil
}
