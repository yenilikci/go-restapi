package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generatePassword(length int) string {
	x := make([]byte, length)
	for i := range x {
		x[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(x)
}

func main() {
	password := generatePassword(15)
	fmt.Println(password) //ex : czyO3VnK2yzcUuQ
}
