package absfactory

import "fmt"

// StripeCreditCard represents Stripe CreditCard Payment
type StripeCreditCard struct{}

func (scc *StripeCreditCard) ProcessPayment(amount int) {
	fmt.Printf("Processed Payment by Stripe Credit Card: ￥ %d", amount)
}

// StripeBankTransfer represents Stripe BankTranfer Payment
type StripeBankTransfer struct{}

func (sbt *StripeBankTransfer) ProcessPayment(amount int) {
	fmt.Printf("Processed Payment by Stripe Bank Transfer: ￥ %d", amount)
}

// StripeFactory creates Stripe Payments Processors
type StripeFactory struct{}

func (sf *StripeFactory) CreditCardProcessor() *StripeCreditCard {
	return &StripeCreditCard{}
}

func (sf *StripeFactory) BankTransferProcessor() *StripeBankTransfer {
	return &StripeBankTransfer{}
}
