package main

import (
	"database/sql"
	"fmt"

	. "./toml/models"            //config.go auto load
	"github.com/BurntSushi/toml" //toml package
	_ "github.com/lib/pq"        //underscore, load, add app domain
)

func main() {

	//Example 1
	var conf ConfigYaml
	if _, err := toml.DecodeFile("./toml/configurations/config.toml", &conf); err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%#v\n", conf.Database)

	//Example 2
	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", conf.Database.User, conf.Database.Password, conf.Database.Server, conf.Database.Database)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT 5 + 5").Scan(&count)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

}
