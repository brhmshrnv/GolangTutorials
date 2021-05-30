package main

import (
	"fmt"
	"strconv"
)

func main() {

	var myString = "1"
	var x = 10
	var f float32 = 2.0

	println(myString, x, f)

	//cast String to Int
	number, _ := strconv.Atoi(myString)
	fmt.Println(number)
	println(number)

	result := number + 2
	println(result)

	//cast Int to String
	strconv.Itoa(result)
	println(result)

	//CASTING
	var i int = 55
	var fl float64 = float64(i)
	//uint is only positive integer
	var uint uint = uint(fl)

	println(i, fl, uint)

}
