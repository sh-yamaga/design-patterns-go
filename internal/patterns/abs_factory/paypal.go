package absfactory

import "fmt"

// payPalCreditCard represents PayPal CreditCard Payment
type payPalCreditCard struct{}

func (ppcc *payPalCreditCard) ProcessPayment(amount int) {
	fmt.Printf("Processed Payment by PayPal Credit Card: ￥ %d\n", amount)
}

// payPalBankTransfer represents PayPal BankTransfer Payment
type payPalBankTransfer struct{}

func (ppbt *payPalBankTransfer) ProcessPayment(amount int) {
	fmt.Printf("Processed Payment by PayPal Credit Card: ￥ %d\n", amount)
}

// PayPalFactory creates PayPal Payments Processors
type PayPalFactory struct{}

func (ppf *PayPalFactory) CreditCardProcessor() *payPalCreditCard {
	return &payPalCreditCard{}
}

func (ppf *PayPalFactory) BankTransferProcessor() *payPalBankTransfer {
	return &payPalBankTransfer{}
}
