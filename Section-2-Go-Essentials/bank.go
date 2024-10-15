package main

import (
	"fmt"
)

func bankApplication() {

	balance := 1000.0
	fmt.Println("Welcome to Go bank")

	for {

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
				continue
			}
			balance += depositAmount

			fmt.Println("Balance upadted: ", balance)
		} else if choice == 3 {
			fmt.Print("Enter the amount you want to withdraw")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > balance {

				fmt.Println("withdraw amount is greater than your balance")
				continue
			}
			balance -= withdrawAmount

			fmt.Println("Balance upadted: ", balance)
		} else {
			break
		}
	}

	fmt.Println("Goodbye")

}

func bankApplicationSwitch() {

	balance := 1000.0
	fmt.Println("Welcome to Go Bank")

	for {
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Balance:", balance)

		case 2:
			fmt.Print("Enter the amount you want to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount")
				continue
			}
			balance += depositAmount
			fmt.Println("Balance updated:", balance)

		case 3:
			fmt.Print("Enter the amount you want to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > balance {
				fmt.Println("Withdraw amount is greater than your balance")
				continue
			}
			balance -= withdrawAmount
			fmt.Println("Balance updated:", balance)

		case 4:
			fmt.Println("Goodbye")
			return // Exit the function

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}

}
