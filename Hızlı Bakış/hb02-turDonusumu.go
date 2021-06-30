package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*
		var myString = "1"
		var x = 10
		var f float32 = 2.0

		fmt.Println(myString, x, f)
	*/

	//string -> int, strconv.Atoi
	var myString = "1"
	number, _ := strconv.Atoi(myString) // _ anlamı blank identifier, err dönerse
	fmt.Println(number)
	result := number + 2
	fmt.Println((result)) //3, convert oldu

	// int -> string, strconv.Itoa
	fmt.Println("Sonuç : " + strconv.Itoa(result)) //Sonuç : 3, convert oldu

	//casting - go dili implicit type c. izin vermez, küçükten büyüğe atarken bilinçli atama yapılmalı
	var i int = 55
	var f float64 = float64(i) //int -> float
	var u uint = uint(f)       //float -> uint (unsigned)
	fmt.Println(i, f, u)       //55 ,55, 55
}
