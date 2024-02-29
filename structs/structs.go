package main

import (
	"fmt"
	"structs/user"
)

func main() {
	userFirstName := getUseData("Enter First Name: ")
	userLastName := getUseData("Enter Last Name: ")
	userBirthDate := getUseData("Enter Birth Date(DD/MM/YYYY): ")

	var appUser *user.User

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthDate)
	if err != nil {
		panic("firstname, lastname, birthdate required!!")
	}
	appUser.OutPutUserData()
	appUser.ClearUserName()
	appUser.OutPutUserData()

	admin := user.NewAdmin("sutanu.das@cloud.com", "12345")
	admin.User.OutPutUserData()
	admin.User.ClearUserName()
	admin.User.OutPutUserData()
}

func getUseData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
