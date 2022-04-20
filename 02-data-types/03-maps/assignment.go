package main

import "fmt"

func main() {
	var command string
	contacts := make(map[string]string)
	fmt.Println("Welcome to your phonebook")

	for {
		fmt.Print("Command> ")
		fmt.Scan(&command)
		if command == "store" {
			fmt.Print("Enter contact: ")
			var contact string
			var no string
			fmt.Scan(&contact, &no)
			contacts[contact] = no
			fmt.Println("Contact saved")
		} else if command == "list" {
			for key, value := range contacts {
				fmt.Println(key, value)
			}
		} else if command == "lookup" {
			fmt.Print("Enter name: ")
			var contact string
			fmt.Scan(&contact)
			fmt.Println(contacts[contact])
		} else if command == "quit" {
			break
		} else {
			fmt.Println("Unknown command: ", command)
		}
	}
	fmt.Println("Bye")
}
