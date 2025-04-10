package absfactory

import "fmt"

// stripeCreditCard represents Stripe CreditCard Payment
type stripeCreditCard struct{}

func (scc *stripeCreditCard) ProcessPayment(amount int) {
	fmt.Printf("Processed Payment by Stripe Credit Card: ￥ %d\n", amount)
}

// stripeBankTransfer represents Stripe BankTranfer Payment
type stripeBankTransfer struct{}

func (sbt *stripeBankTransfer) ProcessPayment(amount int) {
	fmt.Printf("Processed Payment by Stripe Bank Transfer: ￥ %d\n", amount)
}

// StripeFactory creates Stripe Payments Processors
type StripeFactory struct{}

func (sf *StripeFactory) CreditCardProcessor() *stripeCreditCard {
	return &stripeCreditCard{}
}

func (sf *StripeFactory) BankTransferProcessor() *stripeBankTransfer {
	return &stripeBankTransfer{}
}
