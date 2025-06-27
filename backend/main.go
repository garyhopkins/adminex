package main

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func main() {
	// Create a new Huma router (uses chi by default under the hood)
	router := huma.New("My API", "1.0.0")

	// Define the response type
	type Output struct {
		Message string `json:"message"`
	}

	// Register a GET /hello endpoint
	huma.Register(router, huma.Operation{
		Method: "GET",
		Path:   "/hello",
	}, func(ctx context.Context, input struct{}) (*Output, error) {
		return &Output{Message: "Hello, world!"}, nil
	})

	// Start the HTTP server
	http.ListenAndServe(":3000", router)
}

