package main

import (
	"fmt"
	"os"
	"strconv"
)

const maxConcurrency = 3

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := args[0]
	cfg, err := newConfig(rawBaseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("Error while configue %v", err)
	}

	if len(args) == 3 {
		concurrency, err := strconv.Atoi(args[1])
		maxPages, errPages := strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("Error converting concurrency value: %v\n", err)
			os.Exit(1)
		}

		if errPages != nil {
			fmt.Printf("Error converting max pages value: %v\n", err)
			os.Exit(1)
		}
		cfg.concurrencyControl = make(chan struct{}, concurrency)
		cfg.maxPages = maxPages
	}

	if len(args) == 1 {
	}

	fmt.Println("starting crawl of: ", rawBaseURL)
	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d of urls for %s\n", count, normalizedURL)
	}
}
