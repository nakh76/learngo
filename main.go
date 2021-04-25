package main

import (
	"fmt"

	"github.com/nakh76/learngo/mydict"
)

func main() {
	// account := accounts.NewAccount("nakh")
	// account.Deposit(10)
	// err := account.Withdraw((20))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(account)

	dictionary := mydict.Dictionary{"first": "First word"}
	dictionary["hello"] = "hello"

	fmt.Println(dictionary)

}
