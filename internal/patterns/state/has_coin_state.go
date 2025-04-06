package state

import "fmt"

type HasCoinState struct {
	vendingMachine *VendingMachine
}

// InsertCoin
func (hcs *HasCoinState) InsertCoin() {
	fmt.Println("You've already inserted a coin; multiple coins are not needed.")
}

// SelectProduct
func (hcs *HasCoinState) SelectProduct() {
	if hcs.vendingMachine.productCount == 0 {
		fmt.Println("Sorry, the product has been soldout.")
		hcs.vendingMachine.setState(&SoldOutState{
			vendingMachine: hcs.vendingMachine,
		})
	} else {
		fmt.Println("Selected the product.")
		hcs.vendingMachine.setState(&DispendingState{
			vendingMachine: hcs.vendingMachine,
		})
	}
}

// Dispense
func (hcs *HasCoinState) Dispense() {
	fmt.Println("Select Product in advance.")
}
