package payment

import "fmt"

type CreditCard struct{}

func (c *CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Paying with credit card")
	return true
}

type CheckingAccount struct{}

func (c *CheckingAccount) ProcessPayment(amount float32) bool {
	fmt.Println("Paying with checking account")
	return true
}