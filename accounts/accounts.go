package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner string;
	balance int;
}

var errOverdrawn error = errors.New("Bank account is overdrawn");

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{
		owner: owner,
		balance: 0,
	};
	return &account;
}

// Deposit x amount on account
func (a *Account /*set Reciever to make method*/) Deposit(amount int) {
	a.balance += amount;
}

// Balance of the account
func (a Account) Balance() int {
	return a.balance;
}

// Withdraw x amount on account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errOverdrawn;
	}
	a.balance -= amount;
	return nil;
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner;
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner;
}

// Will override default String() of struct
func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s account: ", a.balance);
}