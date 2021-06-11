package main

import "fmt"

func main() {

	myarray := [...]int{131, 23, 42, 3225, 341, 333, 113}
	mySlice := myarray[:]
	fmt.Println(myarray)
	fmt.Println(mySlice)

	mySlice[0] = 11
	fmt.Println(mySlice)
	fmt.Println(myarray)

	//if you change index value in slices , array will change same index value
	//because slice doesnt have own values , gets from array , literally copy from array using pointer in backgorund
	fmt.Println("=====================")

	var colors = []string{"RED", "GREEN", "BLUE"}
	fmt.Println(colors)
	colors = append(colors, "PURPLE")
	fmt.Println(colors)
	fmt.Println("=====================")
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	fmt.Println("=====================")
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)
	fmt.Println("=====================")

}
