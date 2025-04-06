package main

import "github.com/sh-yamaga/design-patterns-go/internal/patterns/state"

func main() {
	// initialize VendingMachine(inventory: 2)
	vm := state.NewVendingMachine(2)

	// purchase 1st product
	vm.InsertCoin()
	vm.SelectProduct()
	vm.Dispense()
}
