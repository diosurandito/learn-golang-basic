package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	names := []string{"Dio", "Bayu", "Odenk", "Gilang"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// For Range

	for i, value := range names {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Dio"
	person["title"] = "Programmer Cupu"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
