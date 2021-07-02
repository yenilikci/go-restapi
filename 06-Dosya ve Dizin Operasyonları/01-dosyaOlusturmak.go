package main

import (
	"fmt"
	"log"
	"os"
)

var (
	newFile *os.File //os içerisindeki File struct'ın pointerı
	err     error    //error tipinde err tanımlamam
)

func main() {

	newFile, err := os.Create("demo.cpp")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(newFile.Name() + " dosyası oluşturuldu")
	}
}
