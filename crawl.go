package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	currentUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error while parsing url %s : %v\n", rawCurrentURL, err)
		return
	}

	if currentUrl.Hostname() != cfg.baseURL.Host {
		return
	}

	normalizedCurrentUrl, errCurrent := normalizeURL(rawCurrentURL)
	if errCurrent != nil {
		return
	}

	isFirst := cfg.addPageVisit(normalizedCurrentUrl)
	if !isFirst {
		return
	}

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("Error HTML ", err)
		return
	}

	urls, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		fmt.Println("Error while getting URLS from HTML ", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
