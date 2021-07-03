package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

var fileFolderPath = "files\\"
var files = []string{fileFolderPath + "demo.go", fileFolderPath + "note1.txt"}

func addFile(fileName string, zw *zip.Writer) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Dosya açılırken bir hata meydana geldi (%s): %s", fileName, err)
	}
	defer file.Close()

	wr, err := zw.Create(fileName)
	if err != nil {
		msg := "%s ZIP dosyası içerisinde yeni bir dosya oluşturulurken hata meydana geldi: %s"
		return fmt.Errorf(msg, fileName, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("%s dosyası ZIP'e yazılırken bir hata meydana geldi: %s", fileName, err)
	}

	return nil

}

func createArchiveZIPFile(archiveFileName string) int {
	if len(archiveFileName) == 0 {
		return -1
	}

	//permission
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(archiveFileName+".zip", flags, 0644)
	if err != nil {
		log.Fatalf("ZIP dosyasına veri yazmak için açılırken bir hata meydana geldi: %s", err)
	}
	defer file.Close()

	zw := zip.NewWriter(file)
	defer zw.Close()

	for _, fileName := range files {
		if err := addFile(fileName, zw); err != nil {
			log.Fatalf("%s dosyasına ZIP dosyasına eklenirken bir hata meydana geldi: %s", fileName, err)
		}
	}
	return 1
}

func main() {
	result := createArchiveZIPFile("zip48833")
	if result > 0 {
		fmt.Println("İşlem başarılı: ", result)
	} else {
		fmt.Println("İşlem başarısız: ", result)
	}
}
