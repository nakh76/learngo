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

	dictionary := mydict.Dictionary{}
	word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// greeting, _ := dictionary.Search(word)
	// fmt.Println(greeting)

	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	dictionary.Add(word, "First")
	greeting, _ := dictionary.Search(word)
	fmt.Println(greeting)

	err := dictionary.Delete("saaa")
	if err != nil {
		fmt.Println(err)
	}
	greeting2, _ := dictionary.Search(word)
	fmt.Println(greeting2)
}
