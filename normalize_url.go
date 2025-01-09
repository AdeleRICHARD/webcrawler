package main

import (
	"errors"
	"net/url"
	"path"
	"strings"

	"golang.org/x/net/html"
)

func normalizeURL(inputUrl string) (string, error) {
	if inputUrl == "" {
		return "", errors.New("InpurUrl is empty")
	}
	parsedURL, err := url.Parse(inputUrl)
	if err != nil {
		return "", err
	}

	cleanedPath := path.Clean(parsedURL.Path)
	if !strings.HasPrefix(cleanedPath, "/") {
		cleanedPath = "/" + cleanedPath
	}

	normalizedURL := parsedURL.Host + cleanedPath
	normalizedURL = strings.ReplaceAll(normalizedURL, " ", "")

	return normalizedURL, nil
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	//tokenizer := html.NewTokenizer(strings.NewReader(htmlBody))
	var urls []string

	/* for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			return nil, tokenizer.Err()
		}

		switch tokenType {
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				urls = append(urls, token.String())
			}
			fmt.Println(token.Data)
		case html.ErrorToken:
			return urls, nil
		}
	} */

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
