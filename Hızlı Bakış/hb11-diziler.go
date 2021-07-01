package main

import "fmt"

func main() {

	myArray1 := [3]int{}
	myArray1[0] = 10
	myArray1[1] = 20
	myArray1[2] = 30
	fmt.Println(myArray1) //[10 20 30]

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)    //[Red Green Blue]
	fmt.Println(colors[1]) //Green"
	colors[1] = "Black"
	fmt.Println(colors[1]) //Black

	var numbers = [5]int{1, 3, 5, 7, 9}
	fmt.Println(numbers)
	fmt.Println("numbers of numbers", len(numbers)) //numbers of numbers 5

	myArray2 := [...]int{3, 5, 7, 9}
	fmt.Println(myArray2) //[3 5 7 9]
	myArray2[3] = 99
	fmt.Println(myArray2) //[3 5 7 99]

	var cars [3]string
	cars[0] = "Tesla"
	cars[1] = "Mercedes"
	cars[2] = "Jaguar"
	fmt.Println(len(cars)) //3
	fmt.Println(cap(cars)) //3

	i := 0
	for i <= len(cars)-1 {
		fmt.Println(cars[i])
		i++
	}
	/*
		Tesla
		Mercedes
		Jaguar
	*/

	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}
	/*
		Tesla
		Mercedes
		Jaguar
	*/

	for i, value := range cars {
		fmt.Println("i = ", i, "& value = ", value)
	}
	/*
		i =  0 & value =  Tesla
		i =  1 & value =  Mercedes
		i =  2 & value =  Jaguar
	*/

	for _, value := range cars {
		fmt.Println(value)
	}
	/*
		Tesla
		Mercedes
		Jaguar
	*/

}
