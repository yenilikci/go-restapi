package main

import "net/http"

func main() {

	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)

}

func search(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")     //key
	rlz := r.FormValue("rlz") //key
	oq := r.FormValue("oq")   //key

	data := "/search?q=" + q + "&rlz=" + rlz + "&oq=" + oq

	w.Write([]byte(data))
}

//google.com/search?q=muhammed+melih+çelik&rlz=1C1FKPE_trTR958TR958&oq=muhammed+melih+çelik
