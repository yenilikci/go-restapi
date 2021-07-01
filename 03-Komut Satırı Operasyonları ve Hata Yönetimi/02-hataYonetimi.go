package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	/*
		file, err := os.Open("dosyam.txt")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(file.Name())
	*/

	myError := errors.New("Bu bir hata!")
	fmt.Println(myError)
	fmt.Println(myError.Error())

	_, err2 := os.Open("abc.rar")
	/*
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	*/
	checkError(err2)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		//örn : log
		log.Fatal()
		os.Exit(1)
	}
}
