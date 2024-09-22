package after

import "fmt"

func (b bank) printBalance() {
	fmt.Printf("Current Balance: %d", b.balance)
}
