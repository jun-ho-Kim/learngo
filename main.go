package main

import (
	"fmt"

	"github.com/junhoKim/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	account.Balance()
	err := account.WithDraw(40)
	if err != nil {
		fmt.Println(err)
	}
	account.Balance()
	account.ChangeOwner("hoho")
	fmt.Println(account)
}
