package main

import "fmt"

func main() {

	myMap := make(map[int]string)
	fmt.Println(myMap) //map[]
	myMap[40] = "Foo"
	myMap[60] = "Bar"
	fmt.Println(myMap) //map[40:Foo 60:Bar]

	states := make(map[string]string)
	states["IST"] = "ISTANBUL"
	states["ANK"] = "ANKARA"
	states["ANT"] = "ANTALYA"
	fmt.Println(states) //map[ANK:ANKARA ANT:ANTALYA IST:ISTANBUL]

	//anahtar adından veri elde et
	antalya := states["ANT"]
	fmt.Println(antalya) //ANTALYA

	//anahtar adından veri sil
	delete(states, "ANK")
	fmt.Println(states) //map[ANT:ANTALYA IST:ISTANBUL]

	//ekle
	states["ERZ"] = "ERZURUM"

	for key, value := range states {
		fmt.Printf("%v: %v\n", key, value)
		/*
			IST: ISTANBUL
			ERZ: ERZURUM
			ANT: ANTALYA
		*/
	}

	//states map nesnesi elemanı sayısınca bir key listesi elde et
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys) //[ERZ ANT IST]

	//key listesindeki indeks değerlerine göre states nesnesinde bulunan şehirler
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
	/*
		ISTANBUL
		ERZURUM
		ANTALYA
	*/

}
