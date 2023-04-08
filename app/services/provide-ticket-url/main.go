package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//current impl for coding session not for production (only prototyping - no logging/error handling)
//programming mode: make it work and then ...

var (
	s3s *s3.Client
	s3p *s3.PresignClient
)

func init() {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	s3s = s3.NewFromConfig(cfg)
	s3p = s3.NewPresignClient(s3s)
}

func handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//fmt.Printf("starting service with %v s3event records.\n", len(s3e.Records))

	s3Key := "msfinder.yaml" //hardcoded for demo

	input := &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("TICKET_BUCKET_NAME")),
		Key:    aws.String(s3Key),
	}

	resp, err := s3p.PresignGetObject(ctx, input, s3.WithPresignExpires(time.Minute*5))
	if err != nil {
		return failed(http.StatusInternalServerError, fmt.Errorf("presigning object from s3 bucket: %w", err))
	}

	url := resp.URL
	fmt.Printf("s3 object %s presign url: %v", s3Key, url)
	fmt.Println("successful processing of the service.")

	response := struct {
		ExportUrl string `json:"exportUrl"`
	}{
		ExportUrl: url,
	}

	return mapResponse(http.StatusOK, response)
}

func failed(statusCode int, err error) (events.APIGatewayProxyResponse, error) {
	return mapResponse(statusCode, err.Error())
}

func mapResponse(statusCode int, response interface{}) (events.APIGatewayProxyResponse, error) {
	body, _ := json.Marshal(response)
	headers := map[string]string{"Content-Type": "application/json"}
	return events.APIGatewayProxyResponse{
		Headers:    headers,
		StatusCode: statusCode,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(handle)
}
