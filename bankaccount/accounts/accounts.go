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
