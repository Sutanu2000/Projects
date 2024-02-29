package main

import (
	"fmt"
	"splitting/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	//we can create the balance.txt file if it is not there at first
	// var accountBalance float64 = 1000
	// fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	//we are calling fileops package here. The error was coming here due to issue in importing path
	//var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		//return
		panic("can't continue")
	}
	fmt.Println("Welcome to the Bank")
	for {
		fmt.Println("What do you want to do")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Balance is:", accountBalance)

		case 2:
			var depositAmount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid ammount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance Updated!! Your acoount balance is:", accountBalance)

		case 3:
			var withdrawAmount float64
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid ammount. Must be greater than 0.")
				continue
			}
			if withdrawAmount >= accountBalance {
				fmt.Println("Invalid ammount. You can't withdraw more than your account balance.")
				continue
			}

			accountBalance -= withdrawAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance Updated!! Your acoount balance is:", accountBalance)

		default:
			fmt.Println("Have a nice day!!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}
