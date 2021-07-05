package main

import (
	"net/http"
)

// / endpoint handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Index Page"))
	x := r.URL.Path[1:]
	///data := "Merhaba " + x
	data := ""
	if len(x) > 0 {
		data = "Path : " + x
	} else {
		data = "Index Page"
	}
	w.Write([]byte(data))
}

// /about endpoint handler
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
	//w.WriteHeader(http.StatusOK)
}

func main() {

	//v1
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello!"))
		})
	*/

	//v2
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index", indexHandler)

	http.HandleFunc("/about", aboutHandler)

	//addr string, handler http.Handler
	http.ListenAndServe(":8080", nil)

}
