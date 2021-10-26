package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fun-facts-fetcher/cmd"
)

func HandleRequest(ctx context.Context) (string, error) {
	awsLambda := cmd.NewAwsLambda()

	return awsLambda.FetchDailyFunFact()
}

func main() {
	lambda.Start(HandleRequest)
}
