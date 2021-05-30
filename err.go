package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	/*file , err :=os.Open("mfile.txt")

	if err!=nil {
		fmt.Println(err.Error())
	}

	fmt.Println(file.Name())*/

	/*
		myError := errors.New("This is error")

		fmt.Println(myError)*/

	_, err := os.Open("abc.rar")
	checkError(err)

}

func checkError(err error) {

	if err != nil {
		fmt.Println("Error : ", err.Error())
		log.Fatal("Error : ", err.Error())
		os.Exit(1)
	}
}
