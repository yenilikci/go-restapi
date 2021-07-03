package main

import "encoding/xml"

func main() {

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
