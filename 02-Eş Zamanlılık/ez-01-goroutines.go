package main

import (
	"runtime"
	"time"
)

/*
	bir goroutines, go tarafından çalışma zamanında yönetilen hafif bir iş parçacığıdır.
	f(x,y,z)
	go f(x,y,z)
*/

func main() {
	//basit bir goroutine kullanımı
	/*
		go xFunc() //abcdefghijklmnopqrstuvwxyz
		time.Sleep(100 * time.Millisecond)
	*/

	runtime.GOMAXPROCS(1)
	go xFuncx()
	time.Sleep(100 * time.Millisecond) //zabcdefghijklmnopqrstuvwxy

}

func xFunc() {
	for l := byte('a'); l <= byte('z'); l++ {
		print(string(l))
	}
}

func xFuncx() {
	for l := byte('a'); l <= byte('z'); l++ {
		go print(string(l))
	}
}
