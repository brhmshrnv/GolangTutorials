package main

import "fmt"

func main() {

	a := 10
	b := 1

	if b > a {
		fmt.Println("Boyukdur")
	} else if b == a {
		fmt.Println("Beraberdir")
	} else {
		fmt.Println("Kicikdir")
	}

	if foo := 2; foo == 1 {
		println("bar")
	} else {
		println("buz")
	}
}
