package main

import (
	"fmt"

	"github.com/Kgl-01/Control_Structures-GO.git/fileOps"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileOps.ReadFloatFromFile(accountBalanceFile, 1000)

	if err != nil {
		fmt.Println("Error: ")
		fmt.Println(err)
		fmt.Println("----------------")
		//panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 @", randomdata.PhoneNumber())

	for {
		presentOptions()
		var choice int
		fmt.Print("Enter your choice: ")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileOps.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 3:
			fmt.Print("Enter amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than zero.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Withdraw successful! Updated balance:", accountBalance)
			fileOps.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Print("Thanks for choosing our bank")
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}
