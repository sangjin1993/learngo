package main

import (
	"fmt"

	"github.com/sangjin1993/learngo/bankaccount/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())

}
