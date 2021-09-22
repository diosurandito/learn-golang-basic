package main

import (
	"fmt"
	"golang-basic/database"
)

// blank identifier gunakan tanda "_"

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
