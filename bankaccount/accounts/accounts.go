package accounts

import "errors"

// Account struct
// struct 이름은 대문자 public 소문자 private
type Account struct {
	// 소문자 private 대문자 public
	owner   string
	balance int
}

// error 발생시 exception
var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
// constructor(생성자)를 만드는 방법
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// receiver 함수이름(매개변수)
// receiver 규칙 struct의 첫 글자를 따서 소문자로 지어야 한다(a Account)
// * pointer receiver
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
// 출금할 때 0원미만이면 에러를 발생하면 error를 리턴
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		// err 발생시 나타나는 문자열 리턴
		return errNoMoney
	}
	a.balance -= amount
	// 아무런 문제가 없으면 nil리턴
	return nil
}
