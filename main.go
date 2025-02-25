package main

import (
	"fmt"
	"os"
	"sort"
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

	fmt.Println("starting crawl of: ", rawBaseURL)
	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	/* for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d of urls for %s\n", count, normalizedURL)
	} */

	printReport(cfg.pages, rawBaseURL)

}

func printReport(pages map[string]int, baseURL string) {
	if baseURL == "" {
		return
	}

	fmt.Println("=============================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	type pageCount struct {
		url   string
		count int
	}

	var sortedPages []pageCount
	for url, count := range pages {
		sortedPages = append(sortedPages, pageCount{url, count})
	}

	sort.Slice(sortedPages, func(i, j int) bool {
		return sortedPages[i].count > sortedPages[j].count
	})

	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.count, page.url)
	}
}
