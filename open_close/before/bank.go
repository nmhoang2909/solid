package after

import (
	"errors"
)

type bank struct {
	balance int
}

func (b bank) deposit(amout int) {
	b.balance += amout
}

func (b bank) withDraw(amout int) error {
	if b.balance < amout {
		return errors.New("Insufficient funds")
	}
	b.balance -= amout
	return nil
}

func (b bank) getBalance() int {
	return b.balance
}
