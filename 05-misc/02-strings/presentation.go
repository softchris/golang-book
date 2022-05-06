package main

import (
	"fmt"
	"strings"
)

type Person struct {
	Name    string
	Address string
	City    string
}

func main() {
	person := Person{Name: "Jean Normand", Address: "123 Way", City: "Washington"}

	fmt.Println(strings.ToLower(person.Name))
	fmt.Println(person.Address)
	fmt.Println(strings.ToUpper(person.City))
}
