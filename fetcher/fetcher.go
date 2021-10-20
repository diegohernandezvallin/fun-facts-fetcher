package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/distillery/aws/fun-facts-fetcher/httpclient"
	"github.com/distillery/aws/fun-facts-fetcher/model"
)

var funFactFetcher FunFactFetcher

type Fetcher interface {
	Fetch(url string) (model.FunFact, error)
}

type FunFactFetcher struct {
	HttpClientHandler httpclient.HttpClient
}

func (funFactFetcher FunFactFetcher) Fetch(url string) (model.FunFact, error) {
	httpClientResponse, err := funFactFetcher.HttpClientHandler.Get(url)
	if err != nil {
		return model.FunFact{}, err
	}

	if httpClientResponse.StatusCode == http.StatusNotFound {
		notFoundErr := fmt.Errorf("requested url not found. url: %s", url)

		return model.FunFact{}, notFoundErr
	}

	var funFact model.FunFact
	err = json.Unmarshal(httpClientResponse.ResponseBody, &funFact)
	if err != nil {
		return model.FunFact{}, err
	}

	return funFact, nil
}

func NewFunFactFetcher(httpClientHandler httpclient.HttpClient) Fetcher {
	if funFactFetcher.HttpClientHandler == nil {
		funFactFetcher = FunFactFetcher{
			HttpClientHandler: httpClientHandler,
		}
	}

	return funFactFetcher
}
