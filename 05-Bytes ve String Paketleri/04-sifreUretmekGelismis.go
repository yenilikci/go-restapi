package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Option struct {
	length    int
	upperCase bool
	lowerCase bool
	numbers   bool
	specials  bool
}

var source = rand.NewSource(time.Now().UnixNano())

//const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charsetLowerCase = "abcdefghijklmnopqrstuvwxyz"
const charsetUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charsetNumbers = "0123456789"
const charsetSpecials = "_[.~]!#+&"

func generatePassword(opt Option) (string, error) {
	x := make([]byte, opt.length)
	charset := "."

	if opt.upperCase {
		charset += charsetUpperCase
	}
	if opt.lowerCase {
		charset += charsetLowerCase
	}
	if opt.specials {
		charset += charsetSpecials
	}
	if opt.numbers {
		charset += charsetNumbers
	}
	if charset == "." {
		return "NON-GENERATED!", errors.New("Herhangi bir karakter seçim yapmak zorundasınız!")
	}
	for i := range x {
		x[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(x), nil
}

func main() {

	password, err := generatePassword(Option{length: 15, upperCase: true, lowerCase: true, specials: true, numbers: true})
	fmt.Println(password) //KFgz!lVxPOl!uMc
	if err != nil {
		fmt.Println(err.Error())
	}

	password2, err2 := generatePassword(Option{length: 15, upperCase: true, lowerCase: false, specials: true, numbers: true})
	fmt.Println(password2) //1_I&3AE4J]7MXJQ

	if err2 != nil {
		fmt.Println(err2.Error())
	}

	password3, err3 := generatePassword(Option{length: 15, upperCase: false, lowerCase: false, specials: false, numbers: false})
	fmt.Println(password3)

	if err3 != nil {
		fmt.Println(err3.Error())
	} /*
		NON-GENERATED!
		Herhangi bir karakter seçim yapmak zorundasınız!
	*/
}
