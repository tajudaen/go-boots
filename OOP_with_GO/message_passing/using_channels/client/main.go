package main

import (
	"fmt"
	"github.com/tajud99n/web_dev/OOP_with_GO/message_passing/using_channels/payment"
)

func main() {
	chargeCh := make(chan float32)
	payment.CreateCreditAccount(chargeCh)
	chargeCh <- 500
	var a string
	fmt.Scanln(& a)
}