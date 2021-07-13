package helpers

import (
	"fmt"
	"log"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
	}
}
