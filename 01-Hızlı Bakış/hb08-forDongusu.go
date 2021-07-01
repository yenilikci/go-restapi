package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("value: ", i)
	}

	/*
		for i := 0; ; i++ {
			fmt.Println("value: ", i)
		}
	*/ //mantık aynı

	/*
		for i := 0; i < 3; {
			fmt.Println("value: ", i)
		}
	*/ //mantık aynı

	for i := 0; i < 3; {
		fmt.Println("value: ", i)
		i++
	}
}
