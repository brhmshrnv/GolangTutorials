package main

type Brand string

const (
	FACEBOOK Brand = "Facebook"
	GOOGLE   Brand = "Google"
)

func printBrand(brand Brand) {
	println(brand)
}

func main() {

	printBrand(FACEBOOK)
}
