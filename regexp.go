package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`d([a-z])o`)

	fmt.Println(regex.MatchString("dio"))
	fmt.Println(regex.MatchString("kio"))
	fmt.Println(regex.MatchString("dIo"))

	// -1 untuk seleksi semuanya
	fmt.Println(regex.FindAllString("dio kio dgi dso dbo dro", 10))
}
