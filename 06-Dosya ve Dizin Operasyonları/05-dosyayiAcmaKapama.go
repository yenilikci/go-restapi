package main

import (
	"log"
	"os"
)

func main() {

	/*
		file, err := os.Open("demo.txt")
		if err != nil {
			log.Fatal(err)
		}
		file.Close() //defer file.Close()
	*/

	// params: path, setting, permission
	/*
		os.O_RDONLY: sadece okuma
		os.O_WRONLY: sadece yazma
		os.O_RDWR: okuma ve yazma yapılabilir
		os.O_APPEND: dosyanın sonuna ekle
		os.O_CREATE: dosya yoksa oluştur
		os.O_TRUNC: açılırken dosyayı kes

		ayarlar birden fazla olarak kullanılabilir
		os.O_CREATE|os.O_APPEND
	*/
	file, err := os.OpenFile("demo.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
