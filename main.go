package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fun-facts-fetcher/cmd"
)

func HandleRequest(ctx context.Context) (string, error) {
	awsLambda := cmd.NewAwsLambda()
	log.Println("Hello World")

	return awsLambda.FetchDailyFunFact()
}

func main() {
	lambda.Start(HandleRequest)
}
