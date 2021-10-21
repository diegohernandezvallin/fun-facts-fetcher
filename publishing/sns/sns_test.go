package sns

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/stretchr/testify/assert"
)

const (
	anyMessage = "Any message"
)

type snsPublisherMock struct {
	snsiface.SNSAPI
	publish func(input *sns.PublishInput) (*sns.PublishOutput, error)
}

func (mock *snsPublisherMock) Publish(input *sns.PublishInput) (*sns.PublishOutput, error) {
	return mock.publish(input)
}

func TestPublishSuccessfully(t *testing.T) {
	mock := snsPublisherMock{
		publish: func(input *sns.PublishInput) (*sns.PublishOutput, error) {
			return &sns.PublishOutput{}, nil
		},
	}
	snsPublisher.snsClient = &mock

	err := snsPublisher.Publish(anyMessage)

	assert.NoError(t, err, "no error was expected")
}

func TestPublishError(t *testing.T) {
	mock := snsPublisherMock{
		publish: func(input *sns.PublishInput) (*sns.PublishOutput, error) {
			return nil, fmt.Errorf("Expected error")
		},
	}
	snsPublisher.snsClient = &mock

	err := snsPublisher.Publish(anyMessage)

	assert.Error(t, err, "error expected")
}
