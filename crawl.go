package main

import (
	"fmt"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	normalizedRawUrl, err := normalizeURL(rawBaseURL)
	if err != nil {
		return
	}

	normalizedCurrentUrl, errCurrent := normalizeURL(rawCurrentURL)
	if errCurrent != nil {
		return
	}

	rawDomainUrl := strings.Split(normalizedRawUrl, "/")[0]
	currentDomainUrl := strings.Split(normalizedCurrentUrl, "/")[0]

	if rawDomainUrl != currentDomainUrl {
		fmt.Println("Other domain stoping this crawl ", currentDomainUrl)
		return
	}

	if count, ok := pages[normalizedCurrentUrl]; ok {
		fmt.Println("Already visited ", normalizedCurrentUrl)
		count++
		return
	}

	pages[normalizedCurrentUrl] = 1

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("Error HTML ", err)
		return
	}

	urls, err := getURLsFromHTML(html, normalizedRawUrl)
	if err != nil {
		fmt.Println("Error while getting URLS from HTML ", err)
		return
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
