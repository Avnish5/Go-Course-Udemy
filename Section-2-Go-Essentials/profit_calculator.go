package main

import "fmt"

func profitCalulator() {
	var revenue float64
	var expense float64
	var taxRate float64

	revenue = takeInput("Revenue: ")
	expense = takeInput("Expense: ")
	taxRate = takeInput("Tax Rate: ")

	ebt, profit, ratio := calulate(expense, revenue, taxRate)

	output(ebt, profit, ratio)
}

func takeInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput

}

func calulate(expense, revenue, taxRate float64) (ebt, profit, ratio float64) {

	ebt = expense - revenue
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return
}

func output(ebt, profit, ratio float64) {

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}
