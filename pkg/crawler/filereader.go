package crawler

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type Match struct {
	FileName   string
	LineNumber int
	Column     int
	Text       string
}

type FileReader struct{}

func getPrettyPrint(text string, match string, index int) string {
	value := strings.Replace(text, match, "\x1b[32m\033[1m"+match+"\033[0m\x1b[2m", -1)

	value = "\x1b[2m" + value + "\x1b[0m"
	return value
}

func (f *FileReader) Get(fileName string, searchPattern string) ([]Match, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	var matches []Match
	for index, line := range text {
		r, err := regexp.Compile(searchPattern)
		match := r.FindString(line)
		if err != nil {
			log.Fatal(err)
		}

		if match != "" {
			column := strings.Index(line, match)
			matches = append(matches, Match{
				FileName:   fileName,
				LineNumber: index + 1,
				Column:     column,
				Text:       getPrettyPrint(line, match, column),
			})
		}
	}

	return matches, nil
}
