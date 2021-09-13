package main

import "fmt"

func main() {
	name := "Dio"
	counter := 0
	increment := func() {
		name := "Coeg"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
