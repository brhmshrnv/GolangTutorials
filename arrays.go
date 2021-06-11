package main

import "fmt"

func main() {

	//simple array

	myarray := [3]int{}
	myarray[0] = 32
	myarray[1] = 13
	myarray[2] = 78

	fmt.Println(myarray)

	fmt.Println("===================")
	//colors array
	var colors [3]string
	colors[0] = "RED"
	colors[1] = "GREEN"
	colors[2] = "BLUE"
	fmt.Println(colors)
	colors[0] = "DARKRED"
	fmt.Println(colors)
	fmt.Println("===================")

	//number array
	var mynumbers = [5]int{11, 22, 33, 44, 55}
	fmt.Println("Number of mynumbers array", len(mynumbers))

	fmt.Println("===================")
	//auto size functionallity

	newArray := [...]int{111, 222, 333, 444, 555}
	fmt.Println(newArray)

	fmt.Println("===================")
	var cars [3]string
	cars[0] = "TESLA"
	cars[1] = "MERCEDES"
	cars[2] = "JAGUAR"
	fmt.Println(len(cars))
	fmt.Println(cap(cars))
	fmt.Println("===================")

	i := 0
	for i <= len(cars)-1 {

		fmt.Println(cars[i])
		i++
	}
	fmt.Println("===================")
	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}
	fmt.Println("===================")
	for i, value := range cars {
		fmt.Println("i:", i, "&value = ", value)
	}
	fmt.Println("===================")

	//blank identifier which doesnt need
	for i, _ := range cars {
		fmt.Println("i:", i, "&value = ", cars[i])
	}
	fmt.Println("===================")
}
