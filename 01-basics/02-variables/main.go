package main

import "fmt"

var (
	players        = 3
	replay         = false
	namePlayerOne  = "chris"
	PI             = 3.14
	customerName   = "Adam"
	accountBalance = 20
)

func main() {
	fmt.Println(players)
	fmt.Println(replay)
	fmt.Println(namePlayerOne)
	fmt.Println(PI)

	fmt.Printf("Customer %s has %d on his bank account", customerName, accountBalance)
}
