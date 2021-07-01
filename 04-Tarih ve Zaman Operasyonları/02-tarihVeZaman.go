package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2021, time.July, 10, 20, 0, 0, 0, time.UTC)
	fmt.Printf("Çalışıyor: %s\n", t) //Çalışıyor: 2021-01-10 20:00:00 +0000 UTC

	now := time.Now()
	fmt.Printf("Anlık Zaman: %s\n", now) //Anlık Zaman: 2021-07-01 18:30:39.2593377 +0300 +03 m=+0.002642601

	fmt.Println("Ay: ", time.Now().Month())
	fmt.Println("Gün: ", time.Now().Day())
	fmt.Println("Haftanın Günü: ", time.Now().Weekday())

	tomorrow := time.Now().AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v,%v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year()) //Tomorrow is Friday, July, 2,2021

	longFormat := "Friday, Friday, 2, 2021"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))

	shortFormat := "7/2/2021"
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))

}
