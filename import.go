package main

import (
	"fmt"
	"golang-basic/helper"
)

func main() {
	helper.SayHello("Dio")
	// helper.sayGoodbye("Dio") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error

}
