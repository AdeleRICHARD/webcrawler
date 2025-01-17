package main

import "strings"

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
		return
	}

	if count, ok := pages[normalizedCurrentUrl]; ok {
		count++
		return
	}

	pages[normalizedCurrentUrl] = 1
}
