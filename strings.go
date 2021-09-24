package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Dio Surandito", "Dio"))
	fmt.Println(strings.Split("Dio Surandito", " "))
	fmt.Println(strings.ToLower("Dio Surandito"))
	fmt.Println(strings.ToUpper("Dio Surandito"))
	fmt.Println(strings.Trim("        Dio Surandito           ", " "))
	fmt.Println(strings.ReplaceAll("Dio Surandito Dio Dio", "Dio", "KJX"))

}
