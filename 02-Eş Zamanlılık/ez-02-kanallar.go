package main

import "fmt"

/*
	Kanallar <- kanal operatörü vasıtası ile değer gönderip alabildiğimiz veri hattıdır.
	goroutinelar arasındaki iletişimi sağlarız.

	ch := make(chan int)
	ch <- v
	v := <-ch
	veri akışı ok yönünde gerçekleşmekte.
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	//iki adet goroutine
	go sum(s[:len(s)], c)
	go sum(s[:len(s)/2], c)
	x, y := <-c, <-c  //goroutineler arasında iletişim sağlandı
	fmt.Println(x, y) //6 21
}
