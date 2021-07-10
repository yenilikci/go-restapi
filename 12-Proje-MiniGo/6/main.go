package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"text/template"
)

type Page struct {
	Title   string
	Author  string
	Header  string
	Content string
	URI     string
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	//string birleştirme işlemi
	var builder bytes.Buffer
	builder.WriteString("exstring1\n")
	builder.WriteString("exstring2\n")
	builder.WriteString("exstring3\n")
	builder.WriteString("exstring4\n")
	builder.WriteString("exstring5\n")
	builder.WriteString("exstring6\n")
	uri := "www.example.com"
	page := Page{
		Title:   "yazı1",
		Author:  "yazar1",
		Header:  "başlık1",
		Content: "içerik1",
		URI:     "http://" + uri,
	}
	t, _ := template.ParseFiles("page.html")
	t.Execute(w, page)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
