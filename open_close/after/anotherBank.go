package after

import "fmt"

type anotherBank struct {
	balance int
}

func (ab anotherBank) printBalance() {
	fmt.Printf("Another Bank Balance: %d", ab.balance)
}
