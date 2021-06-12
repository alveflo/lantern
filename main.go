package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		// TODO: Show help section!
		fmt.Println("Please provide a search pattern")
		return
	}

	log.Println("Search pattern: " + os.Args[1])
}
