# Using maps

A map is a complex data structure that enables you to store things in a key-value fashion. This lets you implement scenarios like phone books, translation dictionaries and more.

## Introduction

This chapter will cover:

- Defining a map.
- Reading the values of map by key but also iterating over it.
- Change the content of a map.

## The use case for a map

You will have scenarios when you code that there are things you need to look up. If you use a dictionary for example you might look up how you can translate a word from English to Spanish or vice versa. In programming, you have similar situations, maybe you want to know what service is run on a certain port for example. There are also databases that are based on the concept of having a unique key that points to a certain value.

How all this is implemented is via map structure. The idea is that you define a key and value and collect all those in a group, a map.

## Creating a map

To create a map in Go, we need to use the following syntax:

```go
map[<key type>]<value type>{ ... entries }
```

Here's an example of creating a map structure that could hold a phone book:

```go
phonebook := map[int]string{ 555123: "Robin Hood", 555404: "Sheriff of Nottingham"}
```

We define a map structure with key type `int` and value type string. Then we assign it a value with `{}`. Each entry is defined according to `<key>: <value>` and separated by a comma. So how do we read a value?

### Create a map with `make()`

Another way to create a map us by using the `make()` function. `make()` returns a initialized map if you give it a type like so:

```go
dictionaryEnSv = make(map[string]string)
```

### Adding entries

To add entries to the map, you need to provide it with a key and value entry like so:

```go
dictionaryEnSv["hello"] = "hej"
```

## Read a value by key

Imagine now that we have these two entries, and you want the value given that you have then entry 555404, how would we do that? We use the square brackets like so `[]`:

```go
phonebook[555404] // "Sheriff of Nottingham"
```

### Check for existing entry

So, you learned that `phonebook[555404]` gives you a value back. What if it doesn't exist? What happens if you give it a key that's not stored in the map is that you get nothing back as a result:

```go
phonebook[888] // this prints as empty in the console
```

There's a better way to check this because accessing an entry with a key returns two values, the value, and a Boolean. The Boolean indicates if this key exists in the map. See this code:

```go
_, exist phonebook[888]
fmt.Println(exist) // false
```

Here you get the value back, but you choose to ignore for this one-time occasion to only focus on `exist`, a Boolean that tells you if the entry exists.

You can even use this construct in an if statement:

```go
if _, exist := phonebook[888] {
  // number exist, call person
}
```

## Iterate over a map

We can iterate over a map with a `for` construct and a `range`. Here's how you can iterate:

```go
for key, value := range phonebook {
  fmt.Println(key, value)
}
```

## Delete an entry

To remove an entry from a map, you can use the `delete()` method. The `delete()` method takes the map and the key to delete as parameters, like so:

```go
delete(phonebook, 555404)
```

## Assignment - build a phone book

Here are your contacts:

```output
Alice 555-123 
Bob 555-124
Jean 555-125
```

```console
Welcome to your phonebook.
Command> store
Enter contact: Rob 555-126
Contact saved
Command> list
Alice 555-123 
Bob 555-124
Jean 555-125
Rob 555-126
Command> lookup
Enter name: Alice
Alice has number: 555-123
```

HINT: you might need to use both a map and a slice.

## Solution

```go
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

```

## Challenge

Right now, there's no error checking. Add a check so that if you look up a contact that doesn't exist, you should get an error message. Here's how it could work:

```console
command> lookup
Enter name: Jane
Contact doesn't exist, do you want to add it? y/n: y
Enter contact: Jane 123
Contact saved
```
