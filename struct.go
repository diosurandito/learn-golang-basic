package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var dio Customer
	dio.Name = "Dio Surandito"
	dio.Address = "Brebes"
	dio.Age = 30

	fmt.Println(dio.Name)
	fmt.Println(dio.Address)
	fmt.Println(dio.Age)
	fmt.Println("------------")

	// Struct Literals
	bayu := Customer{
		Name:    "Bayu Adi Prasetiyo",
		Address: "Tegal",
		Age:     30,
	}
	fmt.Println(bayu)

	odenk := Customer{"Ridho G. Maulana", "Teagl", 30}
	fmt.Println(odenk)
}
