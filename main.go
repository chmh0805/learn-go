package main

import (
	"fmt"
	"learngo/accounts"
	"log"
)

func main() {
	account := accounts.NewAccount("Hyuk");
	account.Deposit(100);
	// fmt.Println(account.Owner(), account.Balance());

	err := account.Withdraw(50);
	if err != nil {
		log.Fatalln(err);
	}

	account.ChangeOwner("Sun");
	fmt.Println(account);
}