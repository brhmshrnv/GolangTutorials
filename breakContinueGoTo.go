package main

import "fmt"

func main() {

	/*i := 0

	for {
		if i >= 3 {
			break
		}
		fmt.Println("i nin deyeri : ", i)
		i++
	}*/

	for i := 0; i < 8; i++ {

		if i%2 == 0 {
			continue
		} else {
			fmt.Println("Output : ", i)
		}
	}
}
