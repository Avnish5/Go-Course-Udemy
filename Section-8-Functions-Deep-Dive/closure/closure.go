package closure

import "fmt"

func PracticeClousre() {
	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func(val int) int {
		return val * 2
	})

	double := createTranformer(2)
	triple := createTranformer(3)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTranformer(factor int) func(int) int {

	return func(number int) int {
		return number * factor
	}
}
