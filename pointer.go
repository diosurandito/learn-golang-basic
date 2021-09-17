package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// Pointer on function
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Brebes", "Jawa Tengah", "Indonesia"}
	address2 := &address1
	// var address5 *Address = &Address{"Brebes", "Jawa Tengah", "Indonesia"}
	address2.City = "Tegal"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	address3 := new(Address)
	address4 := address3

	address4.Country = "Indonesia"

	fmt.Println(address3)
	fmt.Println(address4)

	// pointer on function
	var alamat = Address{
		City:     "Brebes",
		Province: "Jawa Tengah",
		Country:  "",
	}

	// var alamatPointer *Address = &alamat

	ChangeCountryToIndonesia(&alamat)
	// ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)

}
