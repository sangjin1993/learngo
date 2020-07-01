package main

import (
	"fmt"

	"github.com/sangjin1993/bankAccount/banking"
)

func main() {
	account := banking.Account{Owner: "nicolas", Balance: 10000}
	fmt.Println(account)
}
