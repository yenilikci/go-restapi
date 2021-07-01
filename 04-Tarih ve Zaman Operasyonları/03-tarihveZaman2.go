package main

import (
	"fmt"
	"time"
)

func main() {

	x := fmt.Println
	now := time.Now()
	x(now) //2021-07-01 19:16:16.111283 +0300 +03 m=+0.002111001

	x(now.Year())
	x(now.Month())
	x(now.Hour())
	x(now.Minute())
	x(now.Second())
	x(now.Nanosecond())
	x(now.Location())
	x(now.Weekday())

	xTime := time.Date(2021, 10, 11, 20, 23, 0, 0, time.UTC)
	x(xTime.Before(now)) //false
	x(xTime.After(now))  //true
	x(xTime.Equal(now))  //false
	diff := now.Sub((xTime))
	x(diff)               //-2452h0m18.7591657s
	x(diff.Hours())       //-2451.9992291949166
	x(diff.Minutes())     //-147119.953751695
	x(diff.Seconds())     //-8.8264249823168e+06
	x(diff.Nanoseconds()) //-8826424982316800

}
