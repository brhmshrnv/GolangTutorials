package main

import "fmt"

func main() {

	//make(T, n)       slice      slice of type T with length n and capacity n
	//make(T, n, m)    slice      slice of type T with length n and capacity m
	numbers := make([]int, 5, 5)
	numbers[0] = 123
	numbers[1] = 1444
	numbers[2] = 7543
	numbers[3] = 998
	numbers[4] = 546
	fmt.Println(numbers)

	//numbers= append(numbers,333)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))
}
