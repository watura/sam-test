package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// Response for lambda
type Response struct {
	Message string `json:"body"`
}

func handler(ctx context.Context, input Response) (Response, error) {
	msg := fmt.Sprintf("[Invoked] %v", input.Message)
	return Response{msg}, nil
}

func main() {
	lambda.Start(handler)
}
