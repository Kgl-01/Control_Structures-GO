package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What fo you want to do?")
	fmt.Println("1.Check the balance")
	fmt.Println("2.Deposit Money")
	fmt.Println("3.Withdraw Money")
	fmt.Println("4.Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	fmt.Println("Your choice: ", choice)
}
