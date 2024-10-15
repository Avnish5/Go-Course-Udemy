package main

import "fmt"

func profitCalulator() {
	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expense: ")
	fmt.Scan(&expense)

	fmt.Print("Tax Rate:")
	fmt.Scan(&taxRate)

	ebt := expense - revenue
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}
