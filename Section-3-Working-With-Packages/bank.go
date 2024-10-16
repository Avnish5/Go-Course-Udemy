package main

import (
	"example.com/packages-app/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFileName = "balance.txt"

func bankApplication() {

	balance, err := fileops.GetFloatFromFile("balance.txt")
	if err != nil {
		fmt.Println("Error:")
		fmt.Print(err)
		panic("cannot continue")
	}
	fmt.Println("Welcome to Go bank")
	fmt.Println("Reach us at :", randomdata.PhoneNumber())

	for {

		presentOptions()

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

	balance, err := fileops.GetFloatFromFile(accountBalanceFileName)
	if err != nil {
		fmt.Println("Error:")
		fmt.Print(err)
		panic("cannot continue")
	}
	fmt.Println("Welcome to Go Bank")
	fmt.Println("Reach us at :", randomdata.PhoneNumber())

	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalanceFileName, balance)
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
			fileops.WriteFloatToFile(accountBalanceFileName, balance)
			fmt.Println("Balance updated:", balance)

		case 4:
			fmt.Println("Goodbye")
			return // Exit the function

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}

}

func main() {

	bankApplicationSwitch()
}
