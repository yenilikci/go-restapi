package main

import (
	"log"
	"os"
)

func main() {

	//yazma izin test
	file, err := os.OpenFile("demo.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) { //örn: saltokunursa
			log.Println("Hata: Yazma izni reddedildi!")
		}
	}
	file.Close()

	//okuma izin test
	file2, err2 := os.OpenFile("demo.txt", os.O_RDONLY, 0666)
	if err2 != nil {
		if os.IsPermission(err) {
			log.Println("Hata: Okuma izni reddedildi!")
		}
	}
	file2.Close()

}
