package main

import (
	"log"
	"os"
)

//dosyayı yeniden isimlendirmek ve taşımak
func main() {
	originalPath := "demo.txt"
	newPath := "./moved/test.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
