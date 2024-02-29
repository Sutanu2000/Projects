package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}
type Admin struct {
	email    string
	password string
	User     User
}

func (u *User) OutPutUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}
func (u *User) ClearUserName() {
	u.firstName = " "
	u.lastName = " "
}
func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstname, lastname, birthdate required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		//this the way to return the embadded User struct we added here
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "-----",
			createdAt: time.Now(),
		},
	}
}
