package main

import "fmt"

func main() {
	var (
		firstName = "Chris"
		age       = 20
	)
	message := "a message"
	const PI = 3.14

	fmt.Println(firstName, age)
	fmt.Println(message)
	fmt.Println(PI)
}
