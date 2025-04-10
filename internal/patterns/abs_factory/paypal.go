package absfactory

import "fmt"

// PayPalCreditCard represents PayPal CreditCard Payment
type PayPalCreditCard struct{}

func (ppcc *PayPalCreditCard) ProcessPayment(amount int) {
	fmt.Printf("Processed Payment by PayPal Credit Card: ￥ %d", amount)
}

// PayPalBankTransfer represents PayPal BankTransfer Payment
type PayPalBankTransfer struct{}

func (ppbt *PayPalBankTransfer) ProcessPayment(amount int) {
	fmt.Printf("Processed Payment by PayPal Credit Card: ￥ %d", amount)
}

// PayPalFactory creates PayPal Payments Processors
type PayPalFactory struct{}

func (ppf *PayPalFactory) CreditCardProcessor() *PayPalCreditCard {
	return &PayPalCreditCard{}
}

func (ppf *PayPalFactory) BankTransferProcessor() *PayPalBankTransfer {
	return &PayPalBankTransfer{}
}
