package main

import "fmt"

func main() {
	const firstName string = "Dio"
	const lastName = "Surandito"
	const age = 24

	const (
		address = "Brebes"
		country = "Indonesia"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(address)
	fmt.Println(country)
}
