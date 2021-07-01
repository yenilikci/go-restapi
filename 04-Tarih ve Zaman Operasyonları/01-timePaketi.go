package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Şuanki Unix Zamanı: %v\n", time.Now().Unix())
	time.Sleep(2 * time.Second)
	fmt.Printf("Şuanki Unix Zamanı: %v\n", time.Now().Unix())
}
