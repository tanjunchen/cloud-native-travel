package main

import (
	"errors"
	"fmt"
)

// 当你传值给函数或方法时 Go 会复制这些值
// 因此 如果你写的函数需要更改状态 你就需要用指针指向你想要更改的值
//


type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Println("address of balance in test is", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}

type Bitcoin int

type WalletBitcoin struct {
	balance Bitcoin
}

func (w *WalletBitcoin) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *WalletBitcoin) Balance() Bitcoin {
	return w.balance
}

func (w *WalletBitcoin) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (w *WalletBitcoin) WithdrawCheck(amount Bitcoin) error {

	if amount > w.balance {
		return InsufficientFundsError
	}

	w.balance -= amount
	return nil
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

