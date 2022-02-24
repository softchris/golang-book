package main

import "fmt"

// normal struct
type Address struct {
	city   string
	street string
	postal string
}

// embedded with different name
type Person struct {
	name    string
	address Address
}

// embedded with same name
type Employee struct {
	Address
	company string
}

func main() {
	var address Address
	address.city = "London"
	address.street = "Buckingham palace"
	address.postal = "SW1"
	fmt.Println(address.city)

	address2 := Address{"New York", "Central park", "111"}
	address3 := Address{city: "LA", street: "Hollywood Boulevard", postal: "123"}

	fmt.Println(address2.city)
	fmt.Println(address3.city)

	person := Person{
		name: "chris",
		address: Address{
			city: "Stockholm",
		},
	}
	fmt.Println(person.address.city)

	employee := Employee{
		Address: Address{
			city: "LA",
		},
		company: "Microsoft",
	}
	fmt.Println(employee.company)
}
