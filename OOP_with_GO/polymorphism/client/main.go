package main

import (
	"github.com/tajud99n/web_dev/OOP_with_GO/polymorphism/payment"
	"github.com/tajud99n/web_dev/OOP_with_GO/polymorphism/paybroker"
)

type PaymentOption interface {
	ProcessPayment(float32) bool
}

func main() {
	var option PaymentOption

	option = &payment.CreditCard{}

	option.ProcessPayment(500)

	option = &payment.CheckingAccount{}

	option.ProcessPayment(500)

	option = &paybroker.PaymentBrokerAccount{}

	option.ProcessPayment(500)
}