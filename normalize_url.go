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

	newInput := strings.ToLower(inputUrl)
	if !strings.HasPrefix(inputUrl, "https://") && !strings.HasPrefix(inputUrl, "http://") {
		newInput = "https://" + inputUrl
	}

	parsedURL, err := url.Parse(newInput)
	if err != nil {
		return "", err
	}

	if parsedURL.Scheme == "" && parsedURL.Host == "" && parsedURL.Path == "" {
		return "", errors.New("empty parseUrl")
	}

	if strings.Contains(parsedURL.Host, "..") {
		return "", errors.New("malformed URL: Invalid host")
	}

	if strings.ContainsAny(parsedURL.Path, "<>") {
		return "", errors.New("malformed URL: Invalid characters in path")
	}

	// On retourne simplement le hostname pour les URLs qui n'ont que "/"
	if parsedURL.Path == "/" || parsedURL.Path == "//" {
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

func getURLsFromHTML(htmlBody string, rawBaseURL *url.URL) ([]string, error) {
	fmt.Println("Getting url form html with url ", rawBaseURL)
	var urls []string
	var traverseNodes func(*html.Node)

	page, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	traverseNodes = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" && attr.Val != "" {
					href, err := url.Parse(attr.Val)
					if err != nil {
						fmt.Printf("couldn't parse href '%v': %v\n", attr.Val, err)
						continue
					}
					urlFound := rawBaseURL.ResolveReference(href)
					urls = append(urls, urlFound.String())
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNodes(child)
		}

	}
	traverseNodes(page)

	return urls, nil
}
