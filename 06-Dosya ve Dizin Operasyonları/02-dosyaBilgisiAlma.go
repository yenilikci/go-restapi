package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo //os içerisindeki FileInfo struct'ın pointerı
	err      error        //error tipinde err tanımlamam
)

func main() {
	//dosya bilgisini döndürür, eğer dosya yoksa hata döner
	fileInfo, err := os.Stat("demo.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name: ", fileInfo.Name())
	fmt.Println("Size in bytes: ", fileInfo.Size())
	fmt.Println("Permissions: ", fileInfo.Mode())
	fmt.Println("Last Modified: ", fileInfo.ModTime())
	fmt.Println("Is Dictionary: ", fileInfo.IsDir())
	fmt.Println("System Interface Type: ", fileInfo.Sys())
}
