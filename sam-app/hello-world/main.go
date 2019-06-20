package main

import (
	"context"
	"encoding/json"
	"os"

	l "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func main() {
	l.Start(handler)
}

// Response for lambda
type Response struct {
	Message string `json:"body"`
}

func handler(ctx context.Context) (Response, error) {
	// Local な場合はlocal用のconfigを使う
	var lmd *lambda.Lambda
	if os.Getenv("AWS_SAM_LOCAL") == "true" {
		lmd = lambda.New(session.New(localConfig()))
	} else {
		lmd = lambda.New(session.New())
	}

	j, err := json.Marshal(Response{"Hello World"})
	if err != nil {
		return Response{}, err
	}

	input := &lambda.InvokeInput{
		FunctionName:   aws.String("HelloInvokedFunction"),
		Payload:        j,
		InvocationType: aws.String("RequestResponse"),
	}
	resp, err := lmd.Invoke(input)
	if err != nil {
		return Response{}, err
	}
	var r Response
	err = json.Unmarshal(resp.Payload, &r)
	if err != nil {
		return Response{}, err
	}

	return Response{r.Message}, nil
}

func localConfig() *aws.Config {
	c := aws.NewConfig()

	c.DisableEndpointHostPrefix = aws.Bool(true)
	c.Endpoint = aws.String("host.docker.internal:3001")
	c.DisableSSL = aws.Bool(true)

	return c
}
