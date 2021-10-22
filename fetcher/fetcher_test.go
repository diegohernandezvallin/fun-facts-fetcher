package fetcher

import (
	"net/http"
	"testing"

	"github.com/fun-facts-fetcher/model"
	"github.com/stretchr/testify/assert"
)

const (
	okResponseBody       = `{"status":true,"data":{"id":"70","fact":"Smoking will void your Apple warranty.","cat":"tech","hits":"305"}}`
	notFoundResponseBody = "Requested url not found in this server"

	anyUrl = ""
)

type mockHttpClient struct {
	do func(request *http.Request) (model.HttpClientResponse, error)
}

func (mock mockHttpClient) Do(request *http.Request) (model.HttpClientResponse, error) {
	return mock.do(request)
}

func (mock mockHttpClient) Get(url string) (model.HttpClientResponse, error) {
	return mock.do(&http.Request{})
}

func TestFetcherOk(t *testing.T) {
	mock := mockHttpClient{
		do: func(request *http.Request) (model.HttpClientResponse, error) {
			okResponse := model.HttpClientResponse{
				ResponseBody: []byte(okResponseBody),
				StatusCode:   http.StatusOK,
			}

			return okResponse, nil
		},
	}
	funFactFetcher.httpClientHandler = mock

	actual, err := funFactFetcher.Fetch(anyUrl)

	assert.NoError(t, err)
	assert.NotEmpty(t, actual.Data.Fact)
}

func TestFetcherNotFound(t *testing.T) {
	mock := mockHttpClient{
		do: func(request *http.Request) (model.HttpClientResponse, error) {
			notFoundResponse := model.HttpClientResponse{
				ResponseBody: []byte(notFoundResponseBody),
				StatusCode:   http.StatusNotFound,
			}

			return notFoundResponse, nil
		},
	}
	funFactFetcher.httpClientHandler = mock

	actual, err := funFactFetcher.Fetch(anyUrl)

	assert.Error(t, err)
	assert.Empty(t, actual.Data.Fact)
}
