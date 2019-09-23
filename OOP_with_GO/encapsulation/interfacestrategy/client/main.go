package main

import "github.com/tajud99n/web_dev/OOP_with_GO/encapsulation/interfacestrategy/payment"

func main() {
	var option payment.PaymentOption

	option = payment.CreateCreditAccount(
		"John Doe",
		"1111-2222-3333-4444",
		9,
		2021,
		123)

	option.ProcessPayment(500)

	option = payment.CreateCashAccount()
	option.ProcessPayment(500)
}