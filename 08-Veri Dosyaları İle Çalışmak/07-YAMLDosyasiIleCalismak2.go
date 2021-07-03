package main

import (
	//m "./models/"
	"io/ioutil"
	//. "./models" //auto load
)

func main() {
	//os.Args[1]
	//go run file.go config.yaml
	fileName := "./config.yaml"
	//var config Config
	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	println(string(source))
	//yaml.Unmarshal(...)
}
