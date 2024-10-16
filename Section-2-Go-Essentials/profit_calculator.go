package main

import (
	"errors"
	"fmt"
	"os"
)

func storeResult(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\n Ratio: %.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func profitCalulator() {

	revenue, err1 := takeInput("Revenue: ")
	expense, err2 := takeInput("Expense: ")
	taxRate, err3 := takeInput("Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Print(err1)
		return
	}

	ebt, profit, ratio := calulate(expense, revenue, taxRate)
	storeResult(ebt, profit, ratio)

	output(ebt, profit, ratio)
}

func takeInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("do not enter negative or 0 value")
	}
	return userInput, nil

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
