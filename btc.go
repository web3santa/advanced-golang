package main

import "errors"

type BitcoinAccount struct {
	balance int
	fee     int
}

func NewBitcoin() *BitcoinAccount {
	return &BitcoinAccount{
		balance: 0,
		fee:     300,
	}
}

func (b *BitcoinAccount) GetBalance() int {
	return b.balance
}

func (b *BitcoinAccount) Deposit(amount int) {
	b.balance += amount
}

func (b *BitcoinAccount) Withraw(amount int) error {
	newBalance := b.balance - amount - b.fee
	if newBalance < 0 {
		return errors.New("Insufficient fund")
	}
	b.balance = newBalance
	return nil
}
