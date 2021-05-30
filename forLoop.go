package main

import "fmt"

func main() {

	//for

	for i := 0; i < 5; i++ {
		fmt.Println("Value: ", i)
	}

	sum := 1

	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}
}
