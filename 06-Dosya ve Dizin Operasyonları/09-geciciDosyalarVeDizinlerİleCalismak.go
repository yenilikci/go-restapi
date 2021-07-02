package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	//geçici klasör oluştur
	tempDirPath, err := ioutil.TempDir("", "geciciKlasor")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Geçici klasör oluşturuldu: ", tempDirPath)

	//geçici dosya oluştur
	tempFile, err := ioutil.TempFile(tempDirPath, "geciciDosya.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Geçici dosya oluşturuldu: ", tempFile.Name())

	//dosya kapat
	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	//silme
	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%s dosyası silindi.", tempFile.Name())
	}

	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%s klasörü silindi.", tempDirPath)
	}

}
