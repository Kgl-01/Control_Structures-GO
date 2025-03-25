package main

import "fmt"

func main() {

	var accountBalance = 1000.0

	for i := 0; i < 2; i++ {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What fo you want to do?")
		fmt.Println("1. Check the balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is:", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				return
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		} else if choice == 3 {
			fmt.Print("Enter amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than zero.")
				return
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				return
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Withdraw successful! Updated balance:", accountBalance)
		} else if choice == 4 {
			fmt.Print("Goodbye!")
		} else {
			fmt.Println("Invalid input")
		}
	}

}
