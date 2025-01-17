package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("error while calling website with status code : %v", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "text/html" {
		return "", errors.New("no html found in response")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	println(string(body))
	return string(body), nil
}
