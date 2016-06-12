package AcronymServerFetcher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpGetReader interface {
	Read(url string) (page string, err error)
}

var emptyString = ""

type HttpClient struct{}

func NewHttpClient() *HttpClient {
	newHttpClient := new(HttpClient)
	return newHttpClient
}

func (reader *HttpClient) Read(url string) (page string, err error) {
	response, err := http.Get(url)
	if err != nil {
		return emptyString, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		wrongCodeErrorMessage := fmt.Sprintf("Wrong status code: %d. Expected: 200", response.StatusCode)
		wrongCodeError := errors.New(wrongCodeErrorMessage)
		return emptyString, wrongCodeError
	}

	body := response.Body
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return emptyString, err
	}

	return string(bytes), nil
}
