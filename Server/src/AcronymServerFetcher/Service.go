package main

import (
	"errors"
	"log"
	"time"
)

// IntervalService is an interface abstracting primitive timer structure
type IntervalService interface {
	IsRunning() bool
	GetInterval() time.Duration
	Start() error
	Stop() error
}

// FetcherService is a concrete implementation of the IntervalService tailored for Fetch class
type FetcherService struct {
	isRunning bool
	interval  time.Duration

	fetcher Fetcher
	timer   *time.Timer
	logger  *log.Logger
}

// NewFetcherService instantiates new FetchService
func NewFetcherService(fetcher Fetcher, interval time.Duration, logger *log.Logger) *FetcherService {
	service := new(FetcherService)
	service.interval = interval
	service.fetcher = fetcher
	service.logger = logger

	return service
}

// IsRunning tells if FetcherService is running
func (f *FetcherService) IsRunning() bool {
	return f.isRunning
}

// GetInterval gets currently used interval by the service
func (f *FetcherService) GetInterval() time.Duration {
	return f.interval
}

// Start func starts the service
// Will return error if service is already started
func (f *FetcherService) Start() error {

	if f.isRunning {
		return errors.New("Service is already running")
	}

	f.logger.Print("[FetcherService.Start] Starting")
	f.timer = time.NewTimer(f.interval)
	f.isRunning = true
	go f.process()
	f.logger.Print("[FetcherService.Start] Started")

	return nil
}

// Stop func stops the service
// Will return error if service is not running
func (f *FetcherService) Stop() error {

	if !f.isRunning {
		return errors.New("Service is not running")
	}

	f.logger.Print("[FetcherService.Stop] Stopping")
	f.isRunning = false
	f.timer.Stop()
	f.logger.Print("[FetcherService.Stop] Stopped")

	return nil
}

func (f *FetcherService) process() {
	for {

		f.logger.Print("[FetcherService.process] Processing started")
		if !f.isRunning {
			return
		}

		err := f.fetcher.UpdateAll()
		if err != nil {
			f.logger.Printf("Failed to UpdateAll with error: %v on %d", err, time.Now().Unix())
		}

		f.logger.Print("[FetcherService.process] Processing completed")

		// Wait for time to fier
		<-f.timer.C
	}
}
