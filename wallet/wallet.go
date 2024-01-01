package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(depositAmount Bitcoin) error {
	if depositAmount > 0 {
		w.balance += depositAmount
		return nil
	} else {
		return ErrIllegalAmountGiven
	}
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
var InsufficientFundsMessageTemplate = "cannot withdraw, insufficient funds. Balance: %s but Withdraw amount: %s"
var ErrIllegalAmountGiven = errors.New("cannot perform the operation, negative value given")

func (w *Wallet) Withdraw(withdrawAmount Bitcoin) error {

	if withdrawAmount > 0 {
		if w.balance < withdrawAmount {
			return ErrInsufficientFunds
			// return fmt.Errorf(InsufficientFundsMessageTemplate, w.balance, withdrawAmount)
		} else {
			w.balance -= withdrawAmount
			return nil
		}
	} else {
		return ErrIllegalAmountGiven
	}
}
