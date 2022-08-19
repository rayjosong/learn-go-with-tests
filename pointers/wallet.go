package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Wallet struct{
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// *Wallet is a pointer to "a" wallet. means same memory
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit() is %v\n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
 	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
