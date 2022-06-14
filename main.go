package main

import (
	"FindAccountById/FindAccount"
	"fmt"
	"log"
)

func main() {
	id, err := FindAccount.FindAccountByID(3)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(id)
}
