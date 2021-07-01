package main

import (
	"bytes"
	"fmt"
)

func main() {

	//temel birleştirme (bytes paketi bağımsız)
	str1 := "abc"
	str2 := "xyz"
	result := str1 + str2
	fmt.Println(result) //abcxyz

	//performanslı buffer
	var x bytes.Buffer
	for i := 0; i < 10; i++ {
		x.WriteString(rastgeleString())
	}
	fmt.Println(x.String()) //$xyz-345$xyz-345$xyz-345$xyz-345$xyz-345$xyz-345$xyz-345$xyz-345$xyz-345$xyz-345
}

func rastgeleString() string {
	return "$xyz-345"
}
