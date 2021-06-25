package crawler

import (
	"os"
	"path/filepath"
)

type Crawler struct{}

func (c *Crawler) Crawl(directory string, searchPattern string) ([]Match, error) {
	var matches []Match
	for _, file := range getFilesInDirectory(directory) {
		filereader := FileReader{}

		results, err := filereader.Get(file, searchPattern)
		if err != nil {
			return nil, err
		}

		if len(results) > 0 {
			for _, res := range results {
				matches = append(matches, res)
			}
		}
	}

	return matches, nil
}

func getFilesInDirectory(directory string) []string {
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}
