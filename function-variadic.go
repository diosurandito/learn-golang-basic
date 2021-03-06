package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 90, 80, 20)
	fmt.Println(total)

	sliceNumbers := []int{10, 90, 80, 20}
	total = sumAll(sliceNumbers...)
	fmt.Println(total)
}
