package main

import (
	"FindAccountById/payments"
	"fmt"
	"log"
)

func main() {
	//id, err := FindAccount.FindAccountByID(3)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	var payment payments.Payment

	p, err := payment.FindPaymentByID("1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("p:", p)

	p1, err := payment.Reject("5")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("p1:", p1)

	//fmt.Println(id)
}
