package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const fetchInterval = time.Duration(6 * time.Hour)
const httpTimeout = time.Duration(10 * time.Second)

func main() {

	// Create files if necessary
	os.Mkdir("logs", 0777)
	os.Mkdir("db", 0777)

	logFileName := fmt.Sprintf("./logs/log-%d.txt", time.Now().Unix())
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "AcronymServerFecher", log.LstdFlags|log.LUTC)
	logger.Print("[.main] Starting application")
	defer logger.Print("[.main] Stopping application")

	reader := NewHttpClient(httpTimeout)
	parser := NewEnglishWikiParser()
	repository := NewSqliteAcronymRepository("./db/AcronymDb.sqlite")
	englishFetcher := NewEnglishFetcher(reader, parser, repository)

	service := NewFetcherService(englishFetcher, fetchInterval, logger)

	err = service.Start()
	if err != nil {
		panic(err)
	}

	// Run indefinitely
	select {}
}
