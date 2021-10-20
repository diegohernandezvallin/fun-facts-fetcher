package model

type HttpClientResponse struct {
	ResponseBody []byte
	StatusCode   int
}

type Data struct {
	Id   string `json:"id"`
	Fact string `json:"fact"`
	Cat  string `json:"cat"`
	Hits string `json:"hits"`
}

type FunFact struct {
	Status bool `json:"status"`
	Data   Data `json:"data"`
}
