package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	zr, err := zip.OpenReader("zip48833.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zr.Close()

	for _, file := range zr.Reader.File {
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)
		//fmt.Println(extractedFilePath)
		/*
			files\demo.go
			files\note1.txt
		*/

		dirName := strings.Split(file.Name, "\\")
		CreateDirIfNotExist(dirName[0])

		if file.FileInfo().IsDir() {
			log.Println("Klasör oluşturuluyor: ", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			log.Println("Dosya çıkarılıyor: ", file.Name)
			outFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outFile.Close()

			_, err2 := io.Copy(outFile, zippedFile)
			if err2 != nil {
				log.Fatal(err2)
			}
		}
	}
}
