# Go from the beginning - Structs

In this chapter, we will learn about structs. A struct is a complex data type capable of holding many fields. It can also be extended to hold behavior.

##  Pre-Lecture Quiz

TODO

## Introduction

This chapter will cover:

- Declaring and inspecting a struct.
- Embedding a struct within another struct.
- Adding implementations to structs.

## Why structs

Lets start with a simple scenario, you have an account balance. You might store it in a variable like so:

```go
accountBalance int32
```

Now that'a great, but if you want to describe something more complex, like a bank account? A bank account consist of more information like that like id, balance, account owner and so on. You could try representing each one of those properties an integers like so:

```go
var accountBalance int32
var owner string
var id int
```

However, what happens if you need to operate on more than one bank account, I mean you could try to store it like so:

```go
var accountBalance int32
var owner string
var id int

var accountBalance2 int32
var owner2 string
var id2 int
```

It doesn't really scale though, what you need is a more complex type, like a `struct` that's able to group all this information like so:

```go
type Account struct {
  accountBalance int32
  owner string
  id int
}
```

## Defining a struct

Ok, so we understand why we need a struct, to gather related information, and we've seen one example so far `Account`. But lets try breaking the parts down and see how we go about defining a struct. Here's what the syntax looks like:

```go
type <a name for the struct> struct {
  ... fields
}
```

Let's show another example but this time we create a struct for an address:

```go
type Address struct {
 city   string
 street string
 postal string
}
```

### Create a struct instance

To create an instance from a struct, we can use one of two approaches:

- **define a variable**, and set the fields after the variable declaration:

   ```go
   var address Address
   address.city = "London"
   address.street = "Buckingham palace"
   address.postal = "SW1"
   ```

- **define all at once**, we can set all the values in one go as well:

   ```go
   address2 := Address{"New York", "Central park", "111"}
   ```

## Embedding a struct

We can also embed a struct in another struct. Lets see we have our `Address` struct, an address is something that a higher level struct like `Person` can use. Here's how that can look:

```go
type Person struct {
 name    string
 address Address
}
```

In this code, the `Person` struct has a field `address` of type `Address`.

To instantiate this struct, we can type like so:

```go
person := Person{
  name: "chris",
  address: Address{
   city: "Stockholm",
  },
 }
```

### Relying on default naming

Note how we created a field `address`, we can skip typing a few characters by defining it like so instead:

```go
type Employee struct {
 Address
 company string
}
```

Note how we omit the name for the field and just type `Address`, this means the field name and field type will be the same name. Creating an instance from it is very similar:

```go
employee := Employee{
  Address: Address{
   city: "LA",
  },
  company: "Microsoft",
 }
```

## Adding implementation to structs

Structs are by its very nature just data fields that describes something complex. You can add behavior to it though by creating functions that operate on a struct. Here's an example:

```go
func (a Address) string() string {
 return fmt.Sprintf("City: %s, Street: %s, Postal address: %s", a.city, a.street, a.postal)
}
```

We've added a `string()` method. The method *belongs* to `Address` and we can see that with `(...)` right after the `func` keyword that takes `a Address`. The rest of the implementation returns a formatted string via `Sprintf()`.  Given the following code:

```go
var address Address
address.city = "London"
address.street = "Buckingham palace"
address.postal = "SW1"
fmt.Println(address.string())
```

We would get the following output, when calling `string()`:

```output
City: London, Street: Buckingham palace, Postal address: SW1
```

## Assignment - defining a struct

Define a struct representing a row in shopping basket for an e-commerce store. 

Here's example output data:

```output
Title, Description, Quantity, Price per unit, Total
LEGO set, 4000 pieces, 1, 600GBP, 600GBP 
```

### Write a program representing the shopping basket

Write a program that iterates over the shopping basked and calculates the total:

```output
Title, Description, Quantity, Price per unit, Total
LEGO set, 4000 pieces, 1, 600GBP, 600GBP
Plushy, plush toy, 3, 5 GBP, 15GBP 

Total: 615 GBP
```

## Solution

Part I

```go
package main

import (
 "fmt"
)

type Row struct {
 Title       string
 Description string
 Quantity    int
 UnitPrice   float32
}

func main() {
 row := Row{
  Title:       "LEGO set",
  Description: "4000 pieces",
  Quantity:    1,
  UnitPrice:   600,
 }
 fmt.Println(row)
}
```

Part II

```go
package main

import (
 "fmt"
)

type Row struct {
 Title       string
 Description string
 Quantity    int
 UnitPrice   float32
}

func main() {
 row := Row{
  Title:       "LEGO set",
  Description: "4000 pieces",
  Quantity:    1,
  UnitPrice:   600,
 }
 row2 := Row{
  Title:       "Plushy",
  Description: "plush toy",
  Quantity:    3,
  UnitPrice:   5,
 }

 basket := make([]Row, 0)
 basket = append(basket, row)
 basket = append(basket, row2)

 var sum int = 0
 for i := 0; i < len(basket); i++ {
  current := basket[i]
  fmt.Println(current)
  sum += current.Quantity * int(current.UnitPrice)
 }
 fmt.Println("Total", sum)
}

```

## Challenge

TODO

##  Pre-Lecture Quiz

TODO

