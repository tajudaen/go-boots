package payment

import (
	"fmt"
)

type PaymentOption interface {
	ProcessPayment(float32) bool
}

type CreditCard struct {
	ownerName string
	cardNumber string
	expirationMonth int
	expirationYear int
	securityCode int
	availableCredit float32
}

func CreateCreditAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName: ownerName,
		cardNumber: cardNumber,
		expirationMonth: expirationMonth,
		expirationYear: expirationYear,
		securityCode: securityCode,
	}
}

func (c *CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Procesing a credit car payment...")
	return true
}

func (c CreditCard) OwnerName() string {
	return c.ownerName
}