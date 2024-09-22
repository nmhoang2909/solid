package after

// extends the functionality by adding specific methods like calculateInterest, which do not alter the core behavior of depositing and withdrawing funds
type savingBank struct {
	bank
}

func (sb savingBank) calculateInterest() {
	// calculate
}
