package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Merhaba Go!")
	fmt.Println("ibrahim.shirinovv@gmail.com")
	fmt.Println("Salam Ibrahim Shirinov")
	fmt.Println("====================")

	uName := os.Getenv("USERNAME")
	uDomain := os.Getenv("USERDOMAIN")
	prosessorArchitecture := os.Getenv("PROCESSOR_ARCHITECTURE")
	prosessorIdentifier := os.Getenv("PROCESSOR_IDENTIFIER")
	processorLevel := os.Getenv("PROCESSOR_LEVEL")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")

	fmt.Println("Username : " + uName)
	fmt.Println("Domain : " + uDomain)
	fmt.Println("Processor Architecture : " + prosessorArchitecture)
	fmt.Println("Processor Identifier : " + prosessorIdentifier)
	fmt.Println("Processor Level : " + processorLevel)
	fmt.Println("Go Root : " + goRoot)
	fmt.Println("Go Path : " + goPath)
	fmt.Println("Home Path : " + homePath)

}
