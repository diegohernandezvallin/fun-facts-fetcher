package cmd

import (
	"fmt"
	"testing"

	"github.com/fun-facts-fetcher/model"
	_ "github.com/fun-facts-fetcher/publishing"
	"github.com/fun-facts-fetcher/repository"
	"github.com/stretchr/testify/assert"
)

const (
	anyFunFact = "You know that any fun fact"
	anyUrl     = "www.anyfun.fact"
)

type mockFunFactFetcher struct {
	fetch func(url string) (model.FunFact, error)
}

func (mock *mockFunFactFetcher) Fetch(url string) (model.FunFact, error) {
	return mock.fetch(url)
}

type mockDynamodbRepository struct {
	putItem func(funFactItem repository.FunFactItem) error
}

func (mock *mockDynamodbRepository) PutItem(funFactItem repository.FunFactItem) error {
	return mock.putItem(funFactItem)
}

type mockSnsPublisher struct {
	publish func(message string) error
}

func (mock *mockSnsPublisher) Publish(message string) error {
	return mock.publish(message)
}

func TestFetchDailyFunFactOk(t *testing.T) {
	funFactFetcherMock := mockFunFactFetcher{
		fetch: func(url string) (model.FunFact, error) {
			return model.FunFact{
				Data: model.Data{
					Fact: anyFunFact,
				},
			}, nil
		},
	}

	dynamodbRespositoryMock := mockDynamodbRepository{
		putItem: func(funFactItem repository.FunFactItem) error {
			return nil
		},
	}

	snsPublisherMock := mockSnsPublisher{
		publish: func(message string) error {
			return nil
		},
	}

	awsLambda := awsLambda{
		snsPublisher:       &snsPublisherMock,
		dynamodbRepository: &dynamodbRespositoryMock,
		funFactFetcher:     &funFactFetcherMock,
	}

	actual, err := awsLambda.FetchDailyFunFact()
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
}

func TestFetchDailyFunFactUrlNotFound(t *testing.T) {
	funFactFetcherMock := mockFunFactFetcher{
		fetch: func(url string) (model.FunFact, error) {
			return model.FunFact{}, fmt.Errorf("requested url nto found")
		},
	}

	dynamodbRespositoryMock := mockDynamodbRepository{
		putItem: func(funFactItem repository.FunFactItem) error {
			return nil
		},
	}

	snsPublisherMock := mockSnsPublisher{
		publish: func(message string) error {
			return nil
		},
	}

	awsLambda := awsLambda{
		snsPublisher:       &snsPublisherMock,
		dynamodbRepository: &dynamodbRespositoryMock,
		funFactFetcher:     &funFactFetcherMock,
	}

	actual, err := awsLambda.FetchDailyFunFact()
	assert.Error(t, err)
	assert.Empty(t, actual)
}
