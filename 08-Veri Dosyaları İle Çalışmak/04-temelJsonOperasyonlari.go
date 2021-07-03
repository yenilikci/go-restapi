package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `
		{
			"data": {
				"object":"card",
				"id":"card_123654789654",
				"firt_name":"Melih",
				"last_name":"Çelik",
				"balance":"55.000"
			}
		}
	`

	//nested map
	var m map[string]map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil { //json -> go object
		panic(err)
	}

	fmt.Println(m) //map[data:map[balance:55.000 firt_name:Melih id:card_123654789654 last_name:Çelik object:card]]

	fmt.Println("-----------------------")

	b, err := json.Marshal(m) //go object -> json
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b)) //{"data":{"balance":"55.000","firt_name":"Melih","id":"card_123654789654","last_name":"Çelik","object":"card"}}
}
