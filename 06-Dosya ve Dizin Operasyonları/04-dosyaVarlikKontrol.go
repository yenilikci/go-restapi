package main

import (
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	fileInfo, err := os.Stat("demo.txt")
	if err != nil {
		if os.IsNotExist(err) { //var mı
			log.Fatal("File does not exists")
		}
	}
	log.Println("File does exist. File information: ")
	log.Println(fileInfo) //yoksa dosya bilgisini bas
}
