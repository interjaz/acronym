package main

import "fmt"

// Fetcher is an interface for abstracting different type of Wikipedia fetchers
type Fetcher interface {
	Language() string
	UpdateAll() error
}

// EnglishFetcher is Fetcher suitable for English wikipedia
type EnglishFetcher struct {
	reader     HttpGetReader
	parser     WikiParser
	repository AcronymRepository
}

// NewEnglishFetcher returns new EnglishFetcher
func NewEnglishFetcher(
	reader HttpGetReader,
	parser WikiParser,
	repository AcronymRepository) *EnglishFetcher {

	fetcher := new(EnglishFetcher)
	fetcher.reader = reader
	fetcher.parser = parser
	fetcher.repository = repository

	return fetcher
}

// Language is a getter which describes which language particular fetcher operates on
func (fetcher *EnglishFetcher) Language() string {
	return "EN"
}

// UpdateAll pefroms fetch operation and updates local store
func (fetcher *EnglishFetcher) UpdateAll() error {

	baseUrl := "https://en.wikipedia.org/wiki/List_of_acronyms"
	urlsParts := []string{
		"",
		":_A", ":_B", ":_C", ":_D", ":_E", ":_F", ":_G", ":_H", ":_I", ":_J",
		":_K", ":_L", ":_M", ":_N", ":_P", ":_O", ":_Q", ":_R", ":_S", ":_T",
		":_U", ":_V", ":_W", ":_X", ":_Y", ":_Z",
	}

	err := fetcher.repository.Open()
	if err != nil {
		return err
	}
	defer fetcher.repository.Close()

	err = fetcher.repository.DeleteAll()
	if err != nil {
		return err
	}

	for _, urlPart := range urlsParts {
		url := baseUrl + urlPart

		page, err := fetcher.reader.Read(url)
		if err != nil {
			return err
		}

		acronyms, _ := fetcher.parser.Parse(page)

		for _, acronym := range acronyms {

			if len(acronym.Acronym) == 0 {
				continue
			}

			dbAcronym := NewAcronym()
			dbAcronym.Acronym = acronym.Acronym
			dbAcronym.Definition = acronym.Description
			dbAcronym.Language = fetcher.Language()
			dbAcronym.Url = acronym.Url

			//fmt.Printf("%v\n", dbAcronym)

			err = fetcher.repository.Insert(dbAcronym)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}
