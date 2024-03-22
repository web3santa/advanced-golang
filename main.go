package main

import "fmt"

type IBankAccount interface {
	GetBalance() int // 100 = 1 dollar
	Deposit(amount int)
	Withraw(amount int) error
}

func main() {
	myAccounts := []IBankAccount{
		NewWallFargo(),
		NewBitcoin(),
	}

	fmt.Println(myAccounts)

	for _, account := range myAccounts {
		account.Deposit(500)
		if err := account.Withraw(400); err != nil {
			fmt.Printf("ERR: %d\n", err)
		}
		balance := account.GetBalance()

		fmt.Printf("balance = %d\n", balance)
	}
}
