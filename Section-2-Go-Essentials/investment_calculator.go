package main

import "fmt"
import "math"

func main() {

	var investmentAmount float64 = 1000
	var expectedReturns = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow((1+expectedReturns/100), years)
	fmt.Print(futureValue)

}
