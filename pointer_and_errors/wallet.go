package main

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of wallet.Deposit is at %v\n", &w.balance)
	w.balance += amount
}