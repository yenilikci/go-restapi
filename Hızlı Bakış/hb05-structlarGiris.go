package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	fmt.Println(Vertex{1, 2}) //{1 2}

	//yapı alanları
	v := Vertex{3, 4}
	v.X = 4
	fmt.Println(v)   //{4 4}
	fmt.Println(v.X) //4

	//yapı göstericileri
	p := &v
	//(*p).X = 1e9
	p.X = 1e9
	fmt.Println(v) //{1000000000 4}

}
