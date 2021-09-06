package main

import "fmt"

func main() {
	var name = "Joko"

	if name == "Dio" {
		fmt.Println("Hello Dio")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hai, Boleh Kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
