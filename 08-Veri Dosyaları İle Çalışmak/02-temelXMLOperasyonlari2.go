package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("sites.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	v := ObjectSites{}
	err = xml.Unmarshal(data, &v) //Deserialization
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
	/*
			{{ sites} 1 [{{ site} www.google.com Arama Motoru Search Engine} {{ site} www.golang.org.tr Go Türkiye Topluluğu Community} {{ site} www.comlaf.com Sosyal İçerik Sitesi Social Network}]
		    <site>
		        <Name>www.google.com</Name>
		        <Description>Arama Motoru</Description>
		        <Category>Search Engine</Category>
		    </site>
		    <site>
		        <Name>www.golang.org.tr</Name>
		        <Description>Go Türkiye Topluluğu</Description>
		        <Category>Community</Category>
		    </site>
		    <site>
		        <Name>www.comlaf.com</Name>
		        <Description>Sosyal İçerik Sitesi</Description>
		        <Category>Social Network</Category>
		    </site>
		}
	*/
}

type ObjectSites struct {
	XMLName     xml.Name `xml:"sites"`
	Version     string   `xml:"version,attr"`
	Sites       []site   `xml:"site"`
	Description string   `xml:",innerxml"`
}

type site struct {
	XMLName     xml.Name `xml:"site"`
	Name        string   `xml:"Name"`        //xml dosyasındaki karşılığı Name
	Description string   `xml:"Description"` //xml dosyasındaki karşılığı Description
	Category    string   `xml:"Category"`    //xml dosyasındaki karşılığı Category
}
