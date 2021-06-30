package main

import "fmt"

//global
var num = 10

// num := 4 func body dışında bu declare çalışmaz

//const aciklama hata

//global const tanımlaması
const pi = 3.14

var random1, random2 = "xdxdxdxd", "sjsjsjsj"

var (
	degisken1 = "deger1"
	degisken2 = "deger2"
)

func main() {

	var message string
	message = "selam go"
	fmt.Println(message)

	var message2 string = "selam"
	fmt.Print(message2)

	var message3 = "merhaba"
	fmt.Print("\n", message3)

	var a, b, c int
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var x, y, z int = 3, 4, 5
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	var message4 = "Hello world"
	var e, r, t = 3, true, 4.5
	fmt.Println(message4, e, r, t)

	/*
		var h int
		var j string
		var l float64
		var i bool
	*/

	/*
		var u int
		var k, o string = "abc", "xyz"
	*/

	/*
		var p = 42
		var s, b = "xyz", true
	*/

	j := 55 //var keyword kullanmadan veya tip belirtmeden, :=
	fmt.Println(j)

	/*
		v, n := "abc", true
		fmt.Println(v, n)

		message5 := "merhaba go"
		fmt.Println(message5)

		l, w, q := 4, "you", false
		fmt.Println(l, w, q)
	*/

	v := "Metin"
	n := 'M' //ascii karşılık
	l := `Metin`
	fmt.Println(v, n, l)

	var myFloat32 float32 = 44.
	myComplex := complex(3, 4)
	println(myFloat32)
	println(myComplex)

	fmt.Println(num)
	fmt.Println(pi)
	fmt.Println(degisken1, degisken2)

}
