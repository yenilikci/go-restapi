package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
}
