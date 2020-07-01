package accounts

// Account struct
// struct 이름은 대문자 public 소문자 private
type Account struct {
	// 소문자 private 대문자 public
	owner   string
	balance int
}

// NewAccount creates Account
// constructor(생성자)를 만드는 방법
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// receiver 함수이름(매개변수)
// receiver 규칙 struct의 첫 글자를 따서 소문자로 지어야 한다(a Account)
func (a Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}
