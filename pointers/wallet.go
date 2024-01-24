package pointers

import (
	"errors"
	"fmt"
)

type Pounds int

func (p Pounds) String() string {
	return fmt.Sprintf("£%d", p)
}

type Wallet struct {
	balance Pounds
}

func (w *Wallet) Deposit(amt Pounds){
	w.balance += amt
}

func (w *Wallet) Withdraw(amt Pounds) error {

	if amt > w.balance {
		return errors.New("insufficient funds")
	}


	w.balance -= amt
	return nil
}

func (w Wallet) Balance() Pounds {
	return w.balance
}
