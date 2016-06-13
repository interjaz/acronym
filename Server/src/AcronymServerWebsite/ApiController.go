package main

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)

type ApiController interface {
	HandleRequest(url string, body interface{}) (response string, err error)
}

type AcronymApiController struct {
	repository AcronymRepository
}

var emptyString = ""
var apiControllerRegex, _ = regexp.Compile(".*/api/v1/([a-zA-Z0-9_]*).*")
var acronymRegex, _ = regexp.Compile(".*/api/v1/Acronym/(.*)")
var acronymRandomRegex, _ = regexp.Compile(".*/api/v1/Random/([0-9]+)")

func NewAcronymApiController(repository AcronymRepository) *AcronymApiController {
	apiHandler := new(AcronymApiController)
	apiHandler.repository = repository

	return apiHandler
}

func (a *AcronymApiController) HandleRequest(url string, body interface{}) (response string, err error) {

	controllerArray := apiControllerRegex.FindStringSubmatch(url)

	if len(controllerArray) != 2 {
		return emptyString, errors.New("Controller not found")
	}

	controller := controllerArray[1]

	switch controller {
	case "Acronym":
		return a.acronymController(url, body)
	case "Random":
		return a.acronymRandomController(url, body)
	}

	return emptyString, errors.New("Controller out of range")
}

func (a *AcronymApiController) acronymController(url string, body interface{}) (response string, err error) {

	err = a.repository.Open()
	if err != nil {
		return emptyString, err
	}
	defer a.repository.Close()

	acronymArray := acronymRegex.FindStringSubmatch(url)
	if len(acronymArray) != 2 {
		return emptyString, errors.New("Not recognized acronym")
	}

	acronyms, err := a.repository.Find(acronymArray[1])
	if err != nil {
		return emptyString, err
	}

	byteResponse, err := json.Marshal(acronyms)
	if err != nil {
		return emptyString, err
	}

	return string(byteResponse), nil
}

func (a *AcronymApiController) acronymRandomController(url string, body interface{}) (response string, err error) {

	err = a.repository.Open()
	if err != nil {
		return emptyString, err
	}
	defer a.repository.Close()

	acronymArray := acronymRandomRegex.FindStringSubmatch(url)
	if len(acronymArray) != 2 {
		return emptyString, errors.New("Not recognized number")
	}

	count, err := strconv.ParseInt(acronymArray[1], 0, 32)
	if err != nil {
		return emptyString, err
	}

	acronyms, err := a.repository.Random(int(count))
	if err != nil {
		return emptyString, err
	}

	byteResponse, err := json.Marshal(acronyms)
	if err != nil {
		return emptyString, err
	}

	return string(byteResponse), nil
}
