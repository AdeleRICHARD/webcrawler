package main

import (
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mutex              *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func newConfig(rawBaseURL string, maxConcurrency int) (*config, error) {
	baseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseUrl,
		mutex:              &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           0,
	}, nil
}

func (cfg *config) addPageVisit(normalizedURL string) bool {
	// lock reading of pages to ensure no concurrency problem
	cfg.mutex.Lock()
	defer cfg.mutex.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}
