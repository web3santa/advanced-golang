package main

import "errors"

type WallsFargo struct {
	balance int
}

func NewWallFargo() *WallsFargo {
	return &WallsFargo{
		balance: 0,
	}
}

func (w *WallsFargo) GetBalance() int {
	return w.balance
}

func (w *WallsFargo) Deposit(amount int) {
	w.balance += amount
}

func (w *WallsFargo) Withraw(amount int) error {
	newBalance := w.balance - amount
	if newBalance < 0 {
		return errors.New("Insufficient fund")
	}
	w.balance = newBalance
	return nil
}
