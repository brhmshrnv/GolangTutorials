package main

import (
	"fmt"
	"time"
)

func main() {

	//create your own date with Date() method

	t := time.Date(2020, time.November, 30, 12, 30, 0, 0, time.UTC)

	//print t as string ,because t type is Time

	fmt.Printf("Go lauch at %s\n", t)

	println("================")

	//get a moment
	now := time.Now()

	//print as string
	fmt.Printf("Time at this moment %s\n", now)

	println("================")

	fmt.Println("Day is : ", t.Day())
	fmt.Println("Month is : ", t.Month())
	fmt.Println("Year is : ", t.Year())

	println("================")

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %s\n", tomorrow)

	println("================")

	tomorrow.Format("January 2, 2006")

}
