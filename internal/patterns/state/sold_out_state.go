package state

import "fmt"

type SoldOutState struct {
	vendingMachine *VendingMachine
}

// InsertCoin
func (sos *SoldOutState) InsertCoin() {
	fmt.Println("Soldout. Cannot insert coins.")
}

// SelectProduct
func (sos *SoldOutState) SelectProduct() {
	fmt.Println("Soldout. Cannnot select any product.")
}

// Dispense
func (sos *SoldOutState) Dispense() {
	fmt.Println("Soldout. Cannot dispense any product.")
}
