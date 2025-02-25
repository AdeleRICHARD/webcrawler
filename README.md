Collecting workspace informationHere's an improved README that matches your Go web crawler project:

# Web Crawler

A concurrent web crawler written in Go that systematically maps internal links within a website's domain.

## Features

- Concurrent crawling with configurable concurrency limit
- Smart URL normalization and deduplication
- Domain-scoped crawling (stays within initial domain)
- Configurable maximum pages limit
- Pretty formatted report output

## Installation

```bash
git clone <your-repository>
cd webcrawler
go build
```

## Usage

```bash
# Basic usage
./crawler <url>

# Advanced usage with concurrency and page limit
./crawler <url> <concurrency> <max-pages>

# Example
./crawler https://blog.boot.dev 3 10
```

### Parameters

- `url`: Starting URL to crawl (required)
- `concurrency`: Number of concurrent crawlers (optional, default: 3)
- `max-pages`: Maximum number of pages to crawl (optional)

## Core Components

- `crawlPage`: Handles concurrent webpage crawling
- `normalizeURL`: URL normalization and validation
- `getURLsFromHTML`: HTML parsing and link extraction
- `config`: Configuration and state management

## Dependencies

- golang.org/x/net: HTML parsing
- github.com/pkg/errors: Error handling
- github.com/stretchr/testify: Testing framework

## Testing

```bash
go test ./...
```

## Example Output

```
=============================
REPORT for https://example.com
=============================
Found 5 internal links to example.com/path1
Found 3 internal links to example.com/path2
Found 1 internal links to example.com/path3
```

## License

MIT License
