package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (string, error) {
	awsLambda := newAwsLambda()

	return awsLambda.fetchDailyFunFact()
}

func main() {
	lambda.Start(HandleRequest)
}
