package practice1

import "fmt"

func Practice1() {

	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumbers(numbers *[]int, transfrom func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transfrom(val))
	}

	return dNumbers
}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}
