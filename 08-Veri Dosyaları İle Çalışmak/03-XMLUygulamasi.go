package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	xmlFile, err := os.Open("Employees.xml")
	if err != nil {
		fmt.Println("Opening file error: ", err)
	}
	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)

	var c Company
	xml.Unmarshal(xmlData, &c)

	fmt.Println(c.Persons)
	/*
		[        ID: 10001 - FirstName: Melih - LastName: Çelik - UserName: yenilikci
		         ID: 10002 - FirstName: Jack - LastName: Jones - UserName: jj
		]
	*/
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id"`
	FirstName string   `xml:"firstName"`
	LastName  string   `xml:"lastName"`
	UserName  string   `xml:"userName"`
}

type Company struct {
	XMLName xml.Name `xml:"company"`
	Persons []Person `xml:"person"`
}

func (p Person) String() string {
	return fmt.Sprintf("\t ID: %d - FirstName: %s - LastName: %s - UserName: %s \n", p.ID, p.FirstName, p.LastName, p.UserName)
}
