package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

//default constructor - varsayılan boş kurucu
func NewHuman() *Human {
	h := new(Human)
	return h
}

//go metot overload desteklemez
func NewHumanWithParams(firstName, lastName string, age int) *Human {
	h := new(Human)
	h.FirstName = firstName
	h.LastName = lastName
	h.Age = age
	return h
}

func main() {

	// V1 - nesne örneği
	/*
		x := Human{FirstName: "Melih"}
		fmt.Println(x.FirstName) //Melih
	*/

	//V2 - new keyword
	/*
		x := new(Human)
		x.FirstName = "Melih"
		fmt.Println(x.FirstName) //Melih
	*/

	//V3 - kurucu metot
	/*
		x := NewHuman()
		x.Age = 21
		fmt.Println(x.Age) //21
	*/

	//V4 - parametreli kurucu metot
	x := NewHumanWithParams("Melih", "Çelik", 21)
	fmt.Println(x.FirstName) //Melih

	//Veri okuma
	var data = x.FirstName + " " + x.LastName + " / " + strconv.Itoa(x.Age)
	fmt.Println(data) //Melih Çelik / 21
}
