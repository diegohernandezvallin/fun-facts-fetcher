package dynamodb

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/fun-facts-fetcher/repository"
)

var dynamodbRepository DynamodbRepository

type DynamodbRepository struct {
	db        dynamodbiface.DynamoDBAPI
	tableName string
}

func NewDynamoRepository(tableName string, db dynamodbiface.DynamoDBAPI) DynamodbRepository {
	if dynamodbRepository.db == nil {
		dynamodbRepository.db = db
		dynamodbRepository.tableName = tableName
	}

	return dynamodbRepository
}

func (dynamodbRepository DynamodbRepository) PutItem(funFactItem repository.FunFactItem) error {
	av, err := dynamodbattribute.MarshalMap(funFactItem)
	if err != nil {
		return fmt.Errorf("unable to marshal new fun fact item: %s", err)
	}

	tableName := dynamodbRepository.tableName

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynamodbRepository.db.PutItem(input)
	if err != nil {
		errMessage := fmt.Errorf("unable to put item. table: %s error: %v", dynamodbRepository.tableName, err)

		return errMessage
	}

	log.Println("Sucessfully put item. TableName: ", dynamodbRepository.tableName)

	return nil
}
