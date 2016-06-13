package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HttpGetReader interface allows for greater abstraction
type HttpGetReader interface {
	Read(url string) (page string, err error)
}

var emptyString = ""

// HttpClient is an implementation of HttpGetReader using httpClient package
type HttpClient struct {
	client *http.Client
}

// HttpClient return new HttpClient
func NewHttpClient(timeout time.Duration) *HttpClient {
	client := new(HttpClient)

	client.client = &http.Client{
		Timeout: timeout,
	}

	return client
}

func (reader *HttpClient) Read(url string) (page string, err error) {

	response, err := reader.client.Get(url)
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
