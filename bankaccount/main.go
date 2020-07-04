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
	err := account.Withdraw(20)
	// error를 보여줄려면 사용해야하는 코드
	// error handle code
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())

	fmt.Println(account)

}
