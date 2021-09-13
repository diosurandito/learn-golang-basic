package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hai", name, "My name is", customer.Name)
}
func main() {
	dio := Customer{
		Name:    "Dio Surandito",
		Address: "Brebes",
		Age:     24,
	}

	dio.sayHi("Bayu")
}
