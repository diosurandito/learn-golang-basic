package main

import "fmt"

func main() {
	name := "Dio"

	switch name {
	case "Dio":
		fmt.Println("Hello Dio")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hai, Kenalan donk...")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
