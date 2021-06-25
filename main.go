package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/alveflo/lantern/pkg/crawler"
)

func main() {
	if len(os.Args) == 1 {
		// TODO: Show help section!
		fmt.Println("Please provide a search pattern")
		return
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	crawler := crawler.Crawler{}
	matches, err := crawler.Crawl(dir, os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	for _, match := range matches {
		fmt.Printf("%s (%d:%d):\n%s\n", match.FileName, match.LineNumber, match.Column, match.Text)
	}
}
