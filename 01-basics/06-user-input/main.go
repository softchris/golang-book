package main

import "fmt"

func main() {
	// var response string
	fmt.Println("hi")

	// // one input
	// fmt.Scan(&response)
	// fmt.Println("User typed ", response)

	// var a1, a2 string

	// // multiple input
	// fmt.Scan(&a1, &a2)

	// // formatted string
	// str := fmt.Sprintf("a1: %s a2: %s", a1, a2)
	// fmt.Println(str)

	var prefix string
	var no int
	// in110
	fmt.Scanf("%3s%d", &prefix, &no)
	fmt.Printf("prefix: %s, invoice no: %d", prefix, no)

}
