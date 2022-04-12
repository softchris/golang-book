package main

import (
	"fmt"
)

func main() {
	// create array
	arr := make([]string, 0)

	var response string
	for {
		fmt.Print("command> ")
		fmt.Scan(&response)
		if response == "quit" {
			break
		} else if response == "new" {
			fmt.Print("Entry:")
			fmt.Scan(&response)
			arr = append(arr, response)
			// save entry to list
			fmt.Println("Saving entry")
		} else if response == "list" {
			// list entries
			fmt.Println("Listing entries")
			for i := 0; i < len(arr); i++ {
				fmt.Println(arr[i])
			}
		} else {
			fmt.Println("Unknown command", response)
		}

	}
	fmt.Println("bye")
}
