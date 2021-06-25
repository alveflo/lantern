package crawler

import (
	"log"
	"testing"
)

func TestSearch(t *testing.T) {
	fr := FileReader{}
	a, err := fr.Get("../../resources/lorem-ipsum.txt", "dolor")

	if err != nil {
		log.Println(err)
	} else {
		log.Println(a)
	}

	if a != nil && err != nil {
		t.Fail()
	}
}
