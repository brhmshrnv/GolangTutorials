package main

import "fmt"

func main() {

	a := [...]string{"a", "b", "c", "d"}

	for i := range a {
		fmt.Println("Array item : ", i, " is ", a[i])
	}

	capitalcity := map[string]string{"Turkey": "Istanbul", "USA": "Washington", "Germany": "Berlin", "France": "Paris"}

	for key, val := range capitalcity {
		fmt.Println("Map item : Capital of", key, " is ", val)
	}
}
