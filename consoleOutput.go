package main

import "fmt"

func main() {

	str1 := "Hello Cow"
	str2 := "Hello Rabbit"
	str3 := "Hello Snake"

	aNumber := 42
	isTrue := true

	stringLength, _ := fmt.Println(str1, str2, str3)

	//if err == null
	println("String length : ", stringLength)
	fmt.Printf("Value of aNumber %v\n", aNumber)
	fmt.Printf("Value of isTrue %v\n", isTrue)
	fmt.Printf("Value of aNumber as float %.2f\n", float64(aNumber))

	//Golang placeHolders

	fmt.Printf("Data types : %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)

	myString := fmt.Sprintf("Data types as var : %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)

	fmt.Print(myString)

}
