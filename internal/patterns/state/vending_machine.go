package state

type VendingMachine struct {
	state        State
	productCount int
}

func NewVendingMachine(productCount int) *VendingMachine {
	vm := &VendingMachine{
		productCount: productCount,
	}

	if productCount > 0 {
		vm.setState(&NoCoinState{
			vendingMachine: vm,
		})
	} else {
		vm.setState(&SoldOutState{
			vendingMachine: vm,
		})
	}

	return vm
}

// setState updates its state
func (vm *VendingMachine) setState(s State) {
	vm.state = s
}

// InsertCoin
func (vm *VendingMachine) InsertCoin() {
	vm.state.InsertCoin()
}

// SelectProduct
func (vm *VendingMachine) SelectProduct() {
	vm.state.SelectProduct()
}

// Dispense
func (vm *VendingMachine) Dispense() {
	vm.state.Dispense()
}
