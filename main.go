package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	if len(args) == 1 {
		fmt.Println("starting crawl of: ", args[0])
	}
	page, err := getHTML(args[0])
	if err != nil {
		fmt.Println("Error while crawling page : ", err)
		os.Exit(1)
	}

	println(page)
}
