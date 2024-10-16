package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appuser user
	appuser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	outputUserData(&appuser)
}
func outputUserData(user *user) {

	//struct are exception you can directly print without derefrence it
	fmt.Println(user.firstName, user.lastName, user.birthdate, user.createdAt)
	//printing using derefrence
	fmt.Println((*user).firstName, (*user).lastName, (*user).birthdate, (*user).createdAt)

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
