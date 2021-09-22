package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument :", args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

	db_username := os.Getenv("DB_USERNAME_1")
	db_password := os.Getenv("DB_PASSWORD_1")
	fmt.Println(db_username)
	fmt.Println(db_password)
}
