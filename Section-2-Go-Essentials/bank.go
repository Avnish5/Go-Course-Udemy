package main

import (
	"fmt"
)

func bankApplication() {

	balance := 1000.0
	fmt.Println("Welcome to Go bank")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {

		fmt.Println("Your Balance: ", balance)
	} else if choice == 2 {
		fmt.Print("Enter the amount you want to deposit")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {

			fmt.Println("Invalid Amount")
			return
		}
		balance += depositAmount

		fmt.Println("Balance upadted: ", balance)
	} else if choice == 3 {
		fmt.Print("Enter the amount you want to withdraw")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount > balance {

			fmt.Println("withdraw amount is greater than your balance")
			return
		}
		balance -= withdrawAmount

		fmt.Println("Balance upadted: ", balance)
	} else {
		fmt.Print("Goodbye")
	}

}
