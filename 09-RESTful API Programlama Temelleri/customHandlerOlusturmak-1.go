package main

import (
	"io"
	"net/http"
)

func main() {
	var i ironman
	var w wolverine

	//http.HandleFunc()
	mux := http.NewServeMux()
	mux.Handle("/ironman", i)   //Mr. Iron!
	mux.Handle("/wolverine", w) //Mr. Wolverine!

	http.ListenAndServe(":8080", mux)
}

type ironman int

//ServeHTTP
func (x ironman) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. Iron!")
}

type wolverine int

//ServeHTTP
func (x wolverine) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. Wolverine!")
}
