package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error:", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR CODE")
	}
	fmt.Println("Aplikasi Running")
}

func main() {
	runApp(true)
	fmt.Println("Tetap jalan bor")

}
