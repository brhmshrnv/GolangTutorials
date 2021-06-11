package main

import "fmt"

func main() {

	//1.
	mymap := make(map[int]string)
	fmt.Println(mymap)
	mymap[43] = "foo"
	mymap[12] = "bar"
	fmt.Println(mymap)

	fmt.Println("================")
	//2.
	states := make(map[string]string)
	states["IST"] = "Istanbul"
	states["BER"] = "Berlin"
	states["WDC"] = "Washington"
	fmt.Println(states)
	fmt.Println("================")
	// get value with key "BER" from city map

	berlin := states["BER"]
	fmt.Println("Selected : ", berlin)
	fmt.Println("================")
	// delete value with key "IST" from city map
	delete(states, "IST")
	fmt.Println(states)

	fmt.Println("================")

	// add to map
	states["BK"] = "BAKU"

	for k, v := range states {

		fmt.Printf("%v: %v\n", k, v)
		//fmt.Println("key: ", k , " value: " , v )
	}
	fmt.Println("================")

	// when add element to state map create new key
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	fmt.Println("================")

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

	fmt.Println("================")
}
