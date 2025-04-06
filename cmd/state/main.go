package main

import "github.com/sh-yamaga/design-patterns-go/internal/patterns/state"

func main() {
	// initialize VendingMachine(inventory: 2)
	vm := state.NewVendingMachine(2)

	// purchase 1st product
	vm.InsertCoin()
	vm.SelectProduct()
	vm.Dispense()
	// Output:
	// Inserted Coin.
	// Selected the product.
	// Dispending...

	// purchase 2nd product
	vm.InsertCoin()
	vm.SelectProduct()
	vm.Dispense()
	// Output:
	// Inserted Coin.
	// Selected the product.
	// Dispending...
	// Soldout.

	// has been soldout
	vm.InsertCoin()
	vm.SelectProduct()
	vm.Dispense()
	// Output:
	// Soldout. Cannot insert coins.
	// Soldout. Cannnot select any product.
	// Soldout. Cannot dispense any product.
}
