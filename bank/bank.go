package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to GO bank")
	fmt.Println("What do you want to do")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. EXIT")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	fmt.Println("Your choice:", choice)

	if choice == 1 {
		fmt.Println("Your balance is:", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else if choice == 3 {
		fmt.Print("Withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	}
}
