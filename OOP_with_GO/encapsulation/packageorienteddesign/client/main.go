package main

import (
	"fmt"

	"github.com/tajud99n/web_dev/OOP_with_GO/encapsulation/packageorienteddesign/payment"
)

func main() {
	credit := payment.CreateCreditAccount(
		"John Doe",
		"1111-2222-3333-4444",
		9,
		2021,
		123,
	)
	fmt.Printf("Owner name: %v\n", credit.OwnerName())
	fmt.Printf("card number: %v\n", credit.CardNumber())
	fmt.Println("Trying to change card number")
	err := credit.SetCardNumber("0000-2222-3333-4444")
	if err != nil {
		fmt.Printf("That didn't work: %v\n", err)
	} else {
		fmt.Printf("card number has been changed to: %v\n", credit.CardNumber())
	}
	fmt.Printf("Available credit: %v\n", credit.AvailableCredit())
}