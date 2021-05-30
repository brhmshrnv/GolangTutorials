package main

func main() {
	/*

		// if your method return 2 different type a normal value and error for ex. you must assing value to your both variable
		but method return only one type . in this case you can ignore variable with underscore _, you should only
		err,value := method()
	*/

	//in this case you never access to  _ values
	var _, x, _ int = 0, 9, 0
	println(x)

}
