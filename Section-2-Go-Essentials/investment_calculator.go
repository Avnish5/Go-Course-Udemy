package main

import "fmt"

func add(a int, b int) int {
	if a < 0 || b < 0 {
		panic("cannot add negative numbers")
	}
	return a + b
}
func main() {

	fmt.Println("Adding 5 and 3:", add(5, 3))

	// This will trigger a panic
	fmt.Println("Adding -1 and 4:", add(-1, 4))

}
