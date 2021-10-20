package fetcher

import (
	"net/http"
	"testing"

	"github.com/fun-facts-fetcher/model"
)

const (
	okResponseBody       = `{"status":true,"data":{"id":"70","fact":"Smoking will void your Apple warranty.","cat":"tech","hits":"305"}}`
	notFoundResponseBody = "Requested url not found in this server"

	anyUrl = ""
)

type okHttpClient struct {
}

func (okHttpClient okHttpClient) Get(url string) (model.HttpClientResponse, error) {
	return okHttpClient.Do(&http.Request{})
}

func (okHttpClient okHttpClient) Do(request *http.Request) (model.HttpClientResponse, error) {
	okResponse := model.HttpClientResponse{
		ResponseBody: []byte(okResponseBody),
		StatusCode:   http.StatusOK,
	}

	return okResponse, nil
}

type notFoundHttpClient struct {
}

func (notFoundHttpClient notFoundHttpClient) Get(url string) (model.HttpClientResponse, error) {
	return notFoundHttpClient.Do(&http.Request{})
}

func (notFoundHttpClient notFoundHttpClient) Do(request *http.Request) (model.HttpClientResponse, error) {
	notFoundResponse := model.HttpClientResponse{
		ResponseBody: []byte(notFoundResponseBody),
		StatusCode:   http.StatusNotFound,
	}

	return notFoundResponse, nil
}

func TestFetcherOk(t *testing.T) {
	var (
		mockHttpClient     okHttpClient
		testFunFactFecther Fetcher

		result model.FunFact
	)

	mockHttpClient = okHttpClient{}
	testFunFactFecther = NewFunFactFetcher(mockHttpClient)

	result, err := testFunFactFecther.Fetch(anyUrl)
	if err != nil {
		t.Fatalf("Expected no error. Got error: %v", err)
	}

	if result.Data.Fact == "" {
		t.Fatalf("Expected fun fact, got an empty one")
	}

	t.Cleanup(teardown)
}

func TestFetcherNotFound(t *testing.T) {
	var (
		mockHttpClient     notFoundHttpClient
		testFunFactFecther Fetcher

		result model.FunFact
	)

	mockHttpClient = notFoundHttpClient{}
	testFunFactFecther = NewFunFactFetcher(mockHttpClient)

	result, err := testFunFactFecther.Fetch(anyUrl)
	if err == nil {
		t.Fatalf("Expected error. No error recieved")
	}

	if result.Data.Fact != "" {
		t.Fatalf("Expected empty fact. Got fact: %s", result.Data.Fact)
	}

	t.Cleanup(teardown)
}

func teardown() {
	funFactFetcher = FunFactFetcher{}
}
