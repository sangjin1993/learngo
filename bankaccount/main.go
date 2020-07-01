package main

import (
	"fmt"

	"github.com/sangjin1993/learngo/bankaccount/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
