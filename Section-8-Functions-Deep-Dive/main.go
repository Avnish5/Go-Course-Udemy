package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3}
	fmt.Println("Sum:", sum(numbers...))     // Output: Sum: 6
	fmt.Println("Sum:", sum(10, 20, 30, 40)) // Output: Sum: 100
	fmt.Println("Sum:", sum())               // Output: Sum: 0
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total

}
