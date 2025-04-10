package main

import (
	"fmt"

	absfactory "github.com/yamaga-shu/design-patterns-go/internal/patterns/abs_factory"
)

func main() {
	// Stripe Payments
	/// Generate StripeFactory
	stripeFactory := &absfactory.StripeFactory{}

	//// Stripe CreditCard Payment
	stripeCard := stripeFactory.CreditCardProcessor()
	stripeCard.ProcessPayment(1000)
	// Output:
	// Processed Payment by Stripe Credit Card: ￥ 1000

	//// Stripe BankTransfer Payment
	stripeBank := stripeFactory.BankTransferProcessor()
	stripeBank.ProcessPayment(1500)
	// Output:
	// Processed Payment by Stripe Bank Transfer: ￥ 1500

	fmt.Println()

	// PayPal Payments
	/// Generate PayPalFactory
	paypalFactory := &absfactory.PayPalFactory{}

	//// PayPal CreditCard Payment
	paypalCard := paypalFactory.CreditCardProcessor()
	paypalCard.ProcessPayment(1200)
	// Output:
	// Processed Payment by PayPal Credit Card: ￥ 1200

	//// PayPal BankTransfer Payment
	paypalBank := paypalFactory.BankTransferProcessor()
	paypalBank.ProcessPayment(800)
	// Output:
	// Processed Payment by PayPal Credit Card: ￥ 800
}
