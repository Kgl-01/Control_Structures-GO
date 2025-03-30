package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = readFloatFromFile(accountBalanceFile, 1000)

	if err != nil {
		fmt.Println("Error: ")
		fmt.Println(err)
		fmt.Println("----------------")
		//panic("Can't continue, sorry.")
	}

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
			writeFloatToFile(accountBalanceFile, accountBalance)
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
			writeFloatToFile(accountBalanceFile, accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Print("Thanks for choosing our bank")
			return
		default:
			fmt.Println("Invalid input")
		}
	}

	fmt.Print("Thanks for choosing our bank")
}

func writeFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
func readFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("Failed to find file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultValue, errors.New("Failed to parse stored value.")
	}

	return value, nil
}
