package httpclient

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/fun-facts-fetcher/model"
	"github.com/fun-facts-fetcher/util"
	"github.com/stretchr/testify/assert"
)

const (
	anyUrl = "www.anyurl.any"

	okResponseBody       = `{"status":true,"data":{"id":"70","fact":"Smoking will void your Apple warranty.","cat":"tech","hits":"305"}}`
	notFoundResponseBody = "Requested url not found in this server"

	funFactId = "70"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestGetRequestOk(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), anyUrl)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       util.StrToReadCloser(okResponseBody),
			Header:     make(http.Header),
		}
	})

	httpClientHandler.Client = client

	response, err := httpClientHandler.Get(anyUrl)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var actual model.FunFact
	json.Unmarshal(response.ResponseBody, &actual)
	assert.Equal(t, funFactId, actual.Data.Id)
	assert.NotEmpty(t, actual.Data.Fact)
}

func TestGetRequestNotFound(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), anyUrl)
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       util.StrToReadCloser(notFoundResponseBody),
			Header:     make(http.Header),
		}
	})

	httpClientHandler.Client = client

	response, err := httpClientHandler.Get(anyUrl)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	var actual model.FunFact
	json.Unmarshal(response.ResponseBody, &actual)
	assert.Empty(t, actual.Data.Id)
}
