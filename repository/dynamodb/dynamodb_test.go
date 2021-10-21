package dynamodb

import (
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/fun-facts-fetcher/repository"
	"github.com/stretchr/testify/assert"
)

const (
	dailyFunFact = "Daily func fact"
)

type dynamodbMockClient struct {
	dynamodbiface.DynamoDBAPI
	putItem func(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}

func (mock *dynamodbMockClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return mock.putItem(input)
}

func TestPutItemInDynamodbSuccessfully(t *testing.T) {
	now := time.Now().Unix()
	funFactItem := repository.FunFactItem{
		LastTimePolled: now,
		FunFact:        dailyFunFact,
	}

	mock := dynamodbMockClient{
		putItem: func(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {

			return &dynamodb.PutItemOutput{}, nil
		},
	}
	dynamodbRepository.db = &mock

	err := dynamodbRepository.PutItem(funFactItem)

	assert.NoError(t, err, nil, "no error was expected")

	t.Cleanup(teardown)
}

func TestPutItemInDynamodbError(t *testing.T) {
	now := time.Now().Unix()
	funFactItem := repository.FunFactItem{
		LastTimePolled: now,
		FunFact:        dailyFunFact,
	}

	mock := dynamodbMockClient{
		putItem: func(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {

			return nil, fmt.Errorf("Expected error")
		},
	}
	dynamodbRepository.db = &mock

	err := dynamodbRepository.PutItem(funFactItem)

	assert.Error(t, err, "error expected")

	t.Cleanup(teardown)
}

func teardown() {
	dynamodbRepository = DynamodbRepository{}
}
