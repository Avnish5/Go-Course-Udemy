package practice1

import "fmt"

func Practice1() {

	array := [7]int{10, 20, 30, 40, 50, 60, 70}

	// Create a slice from the array
	slice := array[1:4] // Slice includes elements at index 1, 2, and 3

	// Display length and capacity
	fmt.Println("Slice:", slice)                  // Output: Slice: [20 30 40]
	fmt.Println("Length of slice:", len(slice))   // Output: Length of slice: 3
	fmt.Println("Capacity of slice:", cap(slice)) // Output: Capacity of slice: 4
}
