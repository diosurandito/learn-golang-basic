package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	sayGoodBye := getGoodBye
	result := sayGoodBye("Dio")
	fmt.Println(result)
	fmt.Println(sayGoodBye("Bayu"))
}
