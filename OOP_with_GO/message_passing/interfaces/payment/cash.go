package payment

import "fmt"

type CashAccount struct{}


func (c *CashAccount) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a cash transaction...")
	return true
}