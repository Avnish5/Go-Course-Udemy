package main

import "fmt"
import "math"

func main() {

	var investmentAmount = 1000
	var expectedReturns = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturns/100), float64(years))
	fmt.Print(futureValue)

}
