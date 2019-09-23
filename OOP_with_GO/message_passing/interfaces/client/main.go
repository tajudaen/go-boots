package main

import (
	"fmt"
	"github.com/tajud99n/web_dev/OOP_with_GO/message_passing/interfaces/payment"
)

func main() {
	var paymentOption payment.PaymentOption

	paymentOption = &payment.CashAccount{}

	ok := paymentOption.ProcessPayment(500)

	paymentOption = &payment.CreditAccount{
		OwnerName: "John Doe",
		CardNumber: "1111-2222-3333-4444",
		ExpirationMonth: 9,
		ExpirationYear: 2021,
		SecurityCode: 123,
	}

	ok = paymentOption.ProcessPayment(500)

	fmt.Println(ok)
}