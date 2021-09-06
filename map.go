package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Dio Surandito",
		"address": "Brebes",
	}
	person["title"] = "Programmer Cupu"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Dio Surandito"
	book["wrong"] = "Ups"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "wrong")
	fmt.Println(book)
	fmt.Println(len(book))
}
