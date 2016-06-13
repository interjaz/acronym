package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// Assure directories exists
	os.Mkdir("log", 0777)

	logFileName := fmt.Sprintf("./log/log-%d.txt", time.Now().Unix())
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	logger := log.New(logFile, "AcronymServerWebsite", log.LstdFlags|log.LUTC)
	logger.Print("[.main] Starting application")
	defer logger.Print("[.main] Closing application")

	dbPath := "../AcronymServerFetcher/db/AcronymDb.sqlite"
	repository := NewSqliteAcronymRepository(dbPath)

	fileHandler := http.FileServer(http.Dir("./web"))

	controller := NewAcronymApiController(repository)
	apiHandler := NewAcronymApiHandler(controller)

	routing := []RouteHandler{
		*NewRouteHandler("/api/v1", apiHandler),
		*NewRouteHandler("/", fileHandler),
	}
	routingHandler := NewRoutingApiHandler(routing, logger)

	http.Handle("/", routingHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
