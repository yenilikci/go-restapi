package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

//object to yaml
func main() {
	p := PersonYaml{"Melih", "Çelik", 21}
	y, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(y))
	/*
		first_name: Melih
		last_name: Çelik
		age: 21
	*/
}

type PersonYaml struct {
	FirstName string `yaml:"first_name"`
	LastName  string `yaml:"last_name"`
	Age       int    `yaml:"age"`
}
