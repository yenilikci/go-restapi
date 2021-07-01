package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Konsol veri çıkışı
	str1 := "the quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog"
	aNumber := 42
	isTrue := true

	stringLength, _ := fmt.Println(str1, str2, str3)

	fmt.Println("String length: ", stringLength) //String length:  49

	fmt.Printf("Value of aNumber: %v\n", aNumber)                     //Value of aNumber: 42
	fmt.Printf("Value of isTrue: %v\n", isTrue)                       //Value of isTrue: true
	fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber)) //Value of aNumber as float: 42.00

	fmt.Printf("Data types: %T, %T, %T, %T and %T\n", str1, str2, str3, aNumber, isTrue) //Data types: string, string, string, int and bool

	myString := fmt.Sprintf("Data types as var: %T, %T, %T, %T and %T", str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString) //Data types as var: string, string, string, int and bool

	//Konsol veri girişi
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter text: ")
	str, _ := reader.ReadString('\n') //string, error -> str, blank identifier
	fmt.Println(str)

	fmt.Print("\nEnter a number: ")
	str, _ = reader.ReadString('\n')                         //dikkat edersek := değil = kullandık
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64) //boşluk temizliği, float'ın 64 biti
	if err != nil {                                          //nil => null demek
		fmt.Println(err)
	} else {
		fmt.Println("Value of number", f)
	}
}
