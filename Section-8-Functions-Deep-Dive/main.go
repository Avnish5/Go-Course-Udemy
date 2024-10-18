package main

import "fmt"

// giving custom type name to a function
type tranformfn func(int) int

func main() {

	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 2, 3}

	//gettting the new array by passing an array and function to the function
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	//getting the function
	transformFun1 := getTransformerFunction(&numbers)
	transformFun2 := getTransformerFunction(&moreNumbers)

	//gettting the new array by passing an array and function to the function
	tranformNumbers := transformNumbers(&numbers, transformFun1)
	tranformMoreNumbers := transformNumbers(&moreNumbers, transformFun2)

	fmt.Println(doubled)
	fmt.Println(tripled)

	fmt.Println(tranformNumbers)
	fmt.Println(tranformMoreNumbers)

}

// it takes array and function as paramter and return the new array
func transformNumbers(numbers *[]int, transfrom tranformfn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transfrom(val))
	}

	return dNumbers
}

// this function returns the function based upon condition
func getTransformerFunction(numbers *[]int) tranformfn {

	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}
