package pointerserrors

import (
	"errors"
	"fmt"
)

// @Title
// @Description
// @Author
// @Update

type Jackcoin int

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (j Jackcoin) String() string {
	return fmt.Sprintf("%d JKC", j)
}

type Wallet struct {
	balance Jackcoin
}

func (w *Wallet) Deposit(coin Jackcoin) {
	w.balance += coin
}

func (w *Wallet) Balance() Jackcoin {
	return w.balance
}

func (w *Wallet) Withdraw(coin Jackcoin) error {
	if coin > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= coin
	return nil
}
