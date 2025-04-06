package state

import "fmt"

type NoCoinState struct {
	vendingMachine *VendingMachine
}

// InsertCoin changes state to HasCoinState from NoCoinState
func (ncs *NoCoinState) InsertCoin() {
	fmt.Println("Inserted Coin.")
	ncs.vendingMachine.setState(&HasCoinState{
		vendingMachine: ncs.vendingMachine,
	})
}

// SelectProduct indicates that the user cannot select products without inserting coins.
func (ncs *NoCoinState) SelectProduct() {
	fmt.Println("You cannot select products without inserting coins.")
}

// Dispense indicates that no product cannnot be dispensed.
func (ncs *NoCoinState) Dispense() {
	fmt.Println("No coins have been inserted, so goods cannot be dispensed.")
}
