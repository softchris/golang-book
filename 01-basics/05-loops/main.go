package main

import (
	"fmt"
)

func main() {
	//
	// for i := 1; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// while
	var keepGoing = true
	answer := ""
	for keepGoing {
		fmt.Println("Type command: ")
		fmt.Scan(&answer)
		if answer == "quit" {
			keepGoing = false
		}
	}
	fmt.Println("program exit")

	// for each
	arr := []string{"arg1", "arg2", "arg3"}
	for i, s := range arr {
		fmt.Printf("index: %d, item: %s \n", i, s)
	}
}
