package main

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"strings"

	"golang.org/x/net/html"
)

func normalizeURL(inputUrl string) (string, error) {
	if inputUrl == "" {
		return "", errors.New("InputUrl is empty")
	}

	parsedURL, err := url.Parse(inputUrl)
	if err != nil {
		return "", err
	}

	// On retourne simplement le hostname pour les URLs qui n'ont que "/"
	if parsedURL.Path == "/" || parsedURL.Path == "" {
		return parsedURL.Host, nil
	}

	cleanedPath := path.Clean(parsedURL.Path)
	if !strings.HasPrefix(cleanedPath, "/") && cleanedPath != "." {
		cleanedPath = "/" + cleanedPath
	}

	normalizedURL := parsedURL.Host
	if cleanedPath != "." {
		normalizedURL += cleanedPath
	}

	normalizedURL = strings.ReplaceAll(normalizedURL, " ", "")
	return normalizedURL, nil
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	fmt.Println("Getting url form html with url ", rawBaseURL)
	var urls []string
	page, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	for node := range page.Descendants() {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" && attr.Val != "" {
					if strings.HasPrefix(attr.Val, "/") {
						urls = append(urls, rawBaseURL+attr.Val)
					} else {
						urls = append(urls, attr.Val)
					}
					break
				}
			}
		}
	}

	return urls, nil
}
