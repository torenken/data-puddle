package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handle(ctx context.Context, s3e events.S3Event) error {
	log.Printf("it works! fine: %v", s3e)
	return nil
}

func main() {
	lambda.Start(handle)
}
