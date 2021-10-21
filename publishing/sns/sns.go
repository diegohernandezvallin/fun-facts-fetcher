package sns

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

var snsPublisher SnsPublisher

type SnsPublisher struct {
	snsClient snsiface.SNSAPI
	topicName string
}

func NewSnsPublisher(snsClient snsiface.SNSAPI, topicName string) SnsPublisher {
	if snsPublisher.snsClient == nil {
		snsPublisher.snsClient = snsClient
		snsPublisher.topicName = topicName
	}

	return snsPublisher
}

func (snsPublisher SnsPublisher) Publish(message string) error {
	input := sns.PublishInput{
		Message:  &message,
		TopicArn: &snsPublisher.topicName,
	}

	_, err := snsPublisher.snsClient.Publish(&input)
	if err != nil {
		return err
	}

	return nil
}
