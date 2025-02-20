package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeURL(t *testing.T) {
	t.Run("remove scheme", func(t *testing.T) {
		expected := "blog.boot.dev/path"
		inputUrl := "https://blog.boot.dev/path"

		got, err := normalizeURL(inputUrl)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("remove / at the end + scheme", func(t *testing.T) {
		expected := "blog.boot.dev/path"
		inputUrl := "https://blog.boot.dev/path/"

		got, err := normalizeURL(inputUrl)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("remove scheme http", func(t *testing.T) {
		expected := "blog.boot.dev/path"
		inputUrl := "http://blog.boot.dev/path"

		got, err := normalizeURL(inputUrl)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("remove scheme http + / at the end", func(t *testing.T) {
		expected := "blog.boot.dev/path"
		inputUrl := "http://blog.boot.dev/path/"

		got, err := normalizeURL(inputUrl)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("remove double /", func(t *testing.T) {
		expected := "blog.boot.dev/path"
		inputUrl := "https://blog.boot.dev//path"

		got, err := normalizeURL(inputUrl)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("remove double dash at the end", func(t *testing.T) {
		expected := "blog.boot.dev/path"
		inputUrl := "https://blog.boot.dev/path//"

		got, err := normalizeURL(inputUrl)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("remove spaces", func(t *testing.T) {
		expected := "blog.boot.dev/path"
		inputUrl := "https://blog.boot.dev/ path"

		got, err := normalizeURL(inputUrl)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("return error if input empty", func(t *testing.T) {
		inputUrl := ""

		_, err := normalizeURL(inputUrl)
		assert.Error(t, err)
	})

	t.Run("return domain only", func(t *testing.T) {
		inputUrl := "https://blog.boot.dev/"
		expected := "blog.boot.dev"

		got, _ := normalizeURL(inputUrl)
		assert.Equal(t, expected, got)
	})

	t.Run("return domain only and not /", func(t *testing.T) {
		inputUrl := "www.boot.dev/"
		expected := "www.boot.dev"

		got, _ := normalizeURL(inputUrl)
		assert.Equal(t, expected, got)
	})

	t.Run("uppercase letter in url", func(t *testing.T) {
		inputUrl := "https://BLOG.boot.dev/PATH"
		expected := "blog.boot.dev/path"

		got, _ := normalizeURL(inputUrl)
		assert.Equal(t, expected, got)
	})

	t.Run("return domain only and not two /", func(t *testing.T) {
		inputUrl := "https://blog.boot.dev//"
		expected := "blog.boot.dev"

		got, _ := normalizeURL(inputUrl)
		assert.Equal(t, expected, got)
	})

	t.Run("bad format url", func(t *testing.T) {
		inputUrl := "https://blog..boot.dev/path"

		_, err := normalizeURL(inputUrl)
		assert.Error(t, err)
	})

}

func TestGetURLsFromHTML(t *testing.T) {
	t.Run("Exact two links from page", func(t *testing.T) {
		inputURLStr := "https://blog.boot.dev"
		inputURL, err := url.Parse(inputURLStr)
		assert.NoError(t, err)

		inputHtml := `
<html>
    <body>
        <a href="/path/one">
            <span>Boot.dev</span>
        </a>
        <a href="https://other.com/path/one">
            <span>Boot.dev</span>
        </a>
    </body>
</html>
`
		expected := []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"}
		got, err := getURLsFromHTML(inputHtml, inputURL)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Exact one link from page", func(t *testing.T) {
		inputURLStr := "https://blog.boot.dev"
		inputURL, err := url.Parse(inputURLStr)
		assert.NoError(t, err)

		inputHtml := `
<html>
    <body>
        <a href="/path/one">
            <span>Boot.dev</span>
        </a>
    </body>
</html>
`
		expected := []string{"https://blog.boot.dev/path/one"}
		got, err := getURLsFromHTML(inputHtml, inputURL)
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Exact 0 link from page", func(t *testing.T) {
		inputURLStr := "https://blog.boot.dev"
		inputURL, err := url.Parse(inputURLStr)
		assert.NoError(t, err)

		inputHtml := `
<html>
    <body>
        <p> no href </p> 
    </body>
</html>
`
		got, err := getURLsFromHTML(inputHtml, inputURL)
		assert.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("Return nil if inputURL empty", func(t *testing.T) {
		inputURL, err := url.Parse("")
		assert.NoError(t, err)

		inputHtml := `
<html>
    <body>
        <a href=""> no href </p> 
    </body>
</html>
`
		got, err := getURLsFromHTML(inputHtml, inputURL)
		assert.NoError(t, err)
		assert.Nil(t, nil, got)
	})

	t.Run("Return nil if no href", func(t *testing.T) {
		inputURLStr := "https://blog.boot.dev"
		inputURL, err := url.Parse(inputURLStr)
		assert.NoError(t, err)

		inputHtml := `
<html>
    <body>
    </body>
</html>
`
		got, err := getURLsFromHTML(inputHtml, inputURL)
		assert.NoError(t, err)
		assert.Nil(t, got)
	})
}
