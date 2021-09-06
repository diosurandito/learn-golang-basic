package main

import "fmt"

func main() {
	var names [3]string
	var values = [3]int{
		90,
		95,
		80,
	}

	names[0] = "Dio"
	names[1] = "Surandito"
	names[2] = "Keren"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println(values)
	//panjang array
	fmt.Println(len(names))
	fmt.Println(len(values))
}
