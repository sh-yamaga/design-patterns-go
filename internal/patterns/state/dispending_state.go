package state

import "fmt"

type DispendingState struct {
	vendingMachine *VendingMachine
}

// InsertCoin
func (ds *DispendingState) InsertCoin() {
	fmt.Println("Dispening. Cannnot insert any coin.")
}

// SelectProduct
func (ds *DispendingState) SelectProduct() {
	fmt.Println("Dispending. Cannnot select any product.")
}

// Dispense
func (ds *DispendingState) Dispense() {
	fmt.Println("Dispending...")

	ds.vendingMachine.productCount--

	if ds.vendingMachine.productCount > 0 {
		ds.vendingMachine.setState(&NoCoinState{
			vendingMachine: ds.vendingMachine,
		})
	} else {
		fmt.Println("Soldout.")
		ds.vendingMachine.setState(&SoldOutState{
			vendingMachine: ds.vendingMachine,
		})
	}
}
