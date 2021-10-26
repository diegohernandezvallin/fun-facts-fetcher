package cmd

import (
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/fun-facts-fetcher/fetcher"
	"github.com/fun-facts-fetcher/httpclient"
	"github.com/fun-facts-fetcher/publishing"
	snsService "github.com/fun-facts-fetcher/publishing/sns"
	"github.com/fun-facts-fetcher/repository"
	dynamoDBRepository "github.com/fun-facts-fetcher/repository/dynamodb"
	"github.com/fun-facts-fetcher/util"
)

const (
	tableNameKey     = "TABLE_NAME"
	funFactApiUrlKey = "FUN_FACT_API_URL"
	snsTopicKey      = "SNS_TOPIC"
)

type awsLambda struct {
	funFactFetcher     fetcher.Fetcher
	dynamodbRepository repository.Repository
	snsPublisher       publishing.Publisher

	funFactApiUrl string
}

func NewAwsLambda() *awsLambda {
	funFactApiUrl := util.Getenv(funFactApiUrlKey, "")
	tableName := util.Getenv(tableNameKey, "")
	snsTopic := util.Getenv(snsTopicKey, "")

	client := http.Client{}
	httpClientHandler := httpclient.NewHttpHandler(&client)
	funFactFetcher := fetcher.NewFunFactFetcher(httpClientHandler)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	dynamoDBClient := dynamodb.New(sess)
	dynamodbRepository := dynamoDBRepository.NewDynamodbRepository(tableName, dynamoDBClient)

	snsClient := sns.New(sess)
	snsPublisher := snsService.NewSnsPublisher(snsClient, snsTopic)

	return &awsLambda{
		funFactFetcher:     funFactFetcher,
		dynamodbRepository: dynamodbRepository,
		snsPublisher:       snsPublisher,
		funFactApiUrl:      funFactApiUrl,
	}
}

func (awsLambda *awsLambda) FetchDailyFunFact() (string, error) {
	funFact, err := awsLambda.funFactFetcher.Fetch(awsLambda.funFactApiUrl)
	if err != nil {
		return "", err
	}

	lastTimePolled := time.Now().Unix()
	funFactItem := repository.FunFactItem{
		LastTimePolled: lastTimePolled,
		FunFact:        funFact.Data.Fact,
	}
	awsLambda.dynamodbRepository.PutItem(funFactItem)

	awsLambda.snsPublisher.Publish(funFact.Data.Fact)

	return funFact.Data.Fact, nil
}
