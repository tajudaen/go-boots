package main

import (
	"github.com/tajud99n/web_dev/OOP_with_GO/composition/type_embedding/payment"
)

func main() {
	ca := &payment.CreditAccount{}
	ca.AvailableFunds()
	ca.ProcessPayment(500)
}