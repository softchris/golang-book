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

func (a *Address) setAddress(copy Address) {
	a.city = copy.city
	a.street = copy.street
	a.postal = copy.postal
}

func (a Address) string() string {
	return fmt.Sprintf("City: %s, Street: %s, Postal address: %s", a.city, a.street, a.postal)
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
	fmt.Println(address.string())

	address2 := Address{"New York", "Central park", "111"}
	address3 := Address{city: "LA", street: "Hollywood Boulevard", postal: "123"}

	fmt.Println(address2.city)

	address2.setAddress(address3)
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

	// employee.company = "test"
	fmt.Println(employee.company)
}
