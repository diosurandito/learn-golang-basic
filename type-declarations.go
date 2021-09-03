package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpDio NoKTP = "127893888888"
	var marriedStatus Married = true

	fmt.Println(noKtpDio)
	fmt.Println(marriedStatus)

}
