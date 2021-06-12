package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		// TODO: Show help section!
		fmt.Println("Please provide a search pattern")
		return
	}

	log.Println("Hello world!")
}
