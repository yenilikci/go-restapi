package main

import "fmt"

//add 2 paramete alır önce değişken adı sonra tipleri int , int döner
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(10, 20))
}
