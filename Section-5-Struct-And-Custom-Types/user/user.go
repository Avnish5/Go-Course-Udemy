package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     User
}

func NewAdmin(email, password string) Admin {

	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "admin",
			lastName:  "yes",
			birthdate: "3",
			createdAt: time.Now(),
		},
	}

}

func (user *User) OutputUserData() {

	fmt.Println(user.firstName, user.lastName, user.birthdate, user.createdAt)

}

func (user *User) ClearUserData() {

	user.firstName = ""
	user.lastName = ""
	user.birthdate = ""

}

func NewUser(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {

		return nil, errors.New("Do not enter empty values")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthDate,
		createdAt: time.Now(),
	}, nil
}
