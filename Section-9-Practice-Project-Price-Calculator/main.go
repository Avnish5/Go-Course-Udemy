package main

import (
	"fmt"

	"example.com/price-calculator-app/cmdmanager"
	// "example.com/price-calculator-app/filemanager"
	"example.com/price-calculator-app/prices"
)

func main() {

	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fn := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}

	}

}

// func main() {
// 	prices := []float64{10, 20, 30}
// 	taxRates := []float64{0, 0.7, 0.1, 0.15}

// 	result := make(map[float64][]float64)

// 	for _, taxRate := range taxRates {

// 		taxIncludedPrices := make([]float64, len(prices))
// 		for priceIndex, price := range prices {
// 			taxIncludedPrices[priceIndex] = price * (1 + taxRate)
// 		}

// 		result[taxRate] = taxIncludedPrices
// 	}

// 	fmt.Println(result)
// }
