package main

import "fmt"

func main() {
	fmt.Println(name)
	fmt.Println(Version)
}

//if variable name start with lower case it will be PRIVATE variable
var name string = "golang"

//if variable name start with upper case it will be PUBLIC variable
var Version string = "1.2.3"

/*

	multiline comment
*/
