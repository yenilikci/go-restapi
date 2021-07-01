package main

import "fmt"

func main() {

	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("i'nin değeri: ", i)
		i++
	}
	/*
		i'nin değeri:  0
		i'nin değeri:  1
		i'nin değeri:  2
	*/

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Çıktı: ", i)
	}
	/*
		Çıktı:  1
		Çıktı:  3
		Çıktı:  5
	*/

	for i := 0; i < 7; i++ {
		if i == 3 {
			continue
		} else if i == 5 {
			break
		}
		fmt.Println("Çıktı: ", i)
	}
	fmt.Println("İşlem devam...")
	/*
		Çıktı:  0
		Çıktı:  1
		Çıktı:  2
		Çıktı:  4
		İşlem devam...
	*/
}
