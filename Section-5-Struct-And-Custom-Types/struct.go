package main

import (
	"example.com/struct-app/user"
	"fmt"
)

type Age int

func (a Age) log() {
	fmt.Print(a)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appuser, err = user.NewUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Print(err)
		return
	}

	var admin = user.NewAdmin("dd", "fd")
	admin.User.OutputUserData()

	appuser.OutputUserData()
	appuser.ClearUserData()
	appuser.OutputUserData()

	var age Age = 5
	age.log()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
