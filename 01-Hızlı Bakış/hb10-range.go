package main

import "fmt"

/*
	for döngüsünün range ("aralık") formu ile bir dilim ("slice") veya eşlem ("map") üzerinde dolaşılır.
	Bir dilim üzerinde dolaşılırken her seferinde iki değer dönülür. İlki indis, ikincisi o indisli elemanın bir kopyası.
*/

func main() {

	//array
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	/*
		Array item 0 is a
		Array item 1 is b
		Array item 2 is c
		Array item 3 is d
	*/

	//map
	baskentler := map[string]string{"Turkey": "Istanbul", "France": "Paris", "Italy": "Roma"}
	for key := range baskentler {
		fmt.Println("Map item: Capital of", key, "is", baskentler[key])
	}
	/*
		Map item: Capital of Turkey is Istanbul
		Map item: Capital of France is Paris
		Map item: Capital of Italy is Roma
	*/

	for key, val := range baskentler {
		fmt.Println("Map item: Capital of", key, "is", val)
	}
	/*
		Map item: Capital of Turkey is Istanbul
		Map item: Capital of France is Paris
		Map item: Capital of Italy is Roma
	*/

}
