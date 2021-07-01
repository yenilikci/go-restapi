package main

import "fmt"

func main() {

	myArray1 := [...]int{15, 25, 35}
	mySlice1 := myArray1[:]
	fmt.Println(mySlice1) //[15 25 35]
	mySlice1[0] = 11
	fmt.Println(mySlice1) //[11 25 35]
	fmt.Println(myArray1) //[11 25 35]

	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors) //[Red Green Blue]
	colors = append(colors, "Purple")
	fmt.Println(colors) //[Red Green Blue Purple]
	colors = append(colors[1:len(colors)])
	fmt.Println(colors) //[Green Blue Purple]
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors) //[Green Blue]

	numbers := make([]int, 5, 5)
	numbers[0] = 123
	numbers[1] = 234
	numbers[2] = 345
	numbers[3] = 456
	numbers[4] = 567
	fmt.Println(numbers) //[123 234 345 456 567]
	numbers = append(numbers, 678)
	fmt.Println(numbers)      //[123 234 345 456 567 678]
	fmt.Println(cap(numbers)) //10
	fmt.Println(len(numbers)) //6
}
