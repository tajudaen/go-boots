package payment

import (
	"fmt"
)

type PaymentOption interface {
	ProcessPayment(float32) bool
}

type CreditAccount struct {
	OwnerName string
	CardNumber string
	ExpirationMonth int
	ExpirationYear int
	SecurityCode int
	availableCredit float32
}

func (c *CreditAccount) ProcessPayment(amount float32) bool {
	fmt.Println("Procesing a credit car payment...")
	return true
}
