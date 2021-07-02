package main

import (
	"io"
	"log"
	"os"
)

func main() {
	originalFile, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	//yeni bir dosya oluştur
	newFile, err := os.Create("./folder/demo_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//kaynaktan hedefe byteları kopyala
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	//dosya içeriği işle, belleği diske boşalt
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

}
