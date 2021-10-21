package httpclient

import (
	"io/ioutil"
	"net/http"

	"github.com/fun-facts-fetcher/model"
)

const (
	ContentTypeHeader = "Content-Type"

	ContentTypeJSON = "application/json"
)

var httpClientHandler HttpClientHandler

type HttpClient interface {
	Get(url string) (model.HttpClientResponse, error)
	Do(request *http.Request) (model.HttpClientResponse, error)
}

type HttpClientHandler struct {
	Client *http.Client
}

func NewHttpHandler(client *http.Client) HttpClientHandler {
	if httpClientHandler.Client == nil {
		httpClientHandler = HttpClientHandler{
			Client: client,
		}
	}

	return httpClientHandler
}

func (httpClientHandler HttpClientHandler) Get(url string) (model.HttpClientResponse, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return model.HttpClientResponse{}, err
	}

	httpClientResponse, err := httpClientHandler.Do(request)
	if err != nil {
		return model.HttpClientResponse{}, err
	}

	return httpClientResponse, nil
}

func (httpClientHandler HttpClientHandler) Do(request *http.Request) (model.HttpClientResponse, error) {
	request.Header.Add(ContentTypeHeader, ContentTypeJSON)

	response, err := httpClientHandler.Client.Do(request)
	if err != nil {
		return model.HttpClientResponse{}, err
	}
	defer response.Body.Close()

	responseBodyParsed, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return model.HttpClientResponse{}, nil
	}

	httpResponse := model.HttpClientResponse{
		ResponseBody: responseBodyParsed,
		StatusCode:   response.StatusCode,
	}

	return httpResponse, nil
}
