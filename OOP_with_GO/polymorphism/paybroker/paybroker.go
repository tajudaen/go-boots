package paybroker

import "fmt"

type PaymentBrokerAccount struct{}

func (c *PaymentBrokerAccount) ProcessPayment(amount float32) bool {
	fmt.Println("Paying with payment broker")
	return true
}