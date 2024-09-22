package after

// goldBank violates by changing the behavior of the deposit method
type goldBank struct {
	bank
}

func (gb goldBank) deposit(amout int) {
	gb.bank.balance += amout
}
