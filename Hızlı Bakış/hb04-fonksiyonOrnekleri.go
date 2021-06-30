package main

import "fmt"

func main() {

	sayHello() //Hello
	//sayHi("Hi") //Hi
	message := "Hi"
	sayHi(message) //Hi
	message = "Hi2"
	sayHi2(message) //Hi2
	//value type
	println(message) //Hi2

	message = "Hi3"
	sayHi3(&message) //Hi3
	//reference type
	println(message) //Hi Go!

	fmt.Println(add(3, 4))      //7
	fmt.Println(add2(3, 4, "")) //7
}

func sayHello() {
	println("Hello")
}

func sayHi(message string) {
	println(message)
}

func sayHi2(message string) {
	println(message)
	message = "Hi Go!"
}

func sayHi3(message *string) {
	println(*message)
	*message = "Hi Go!"
}

// x int, y int veya x,y int
func add(x, y int) int {
	return x + y
}

func add2(x, y int, z string) int {
	return x + y
}
