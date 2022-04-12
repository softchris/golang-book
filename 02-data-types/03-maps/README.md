# Go from the beginning - using maps

A map is a complex data structure that enables you to store things in a key-value fashion. This lets you implement scenarios like phone books, translation dictionaries and more.

##  Pre-Lecture Quiz

TODO

## Introduction

This chapter will cover:

- Defining a map.
- Reading the values of map by key but also iterating over it.
- Change the content of a map.

## The use case for a map

You will have scenarios when you code that there are things you need to look up. If you use a dictionary for example you might look up how you can translate a word from English to Spanish or vice versa. In programming, you have similar situations, maybe you want to know what service is run on a certain port for example. There's also databases that's based around the concept of having a unique key that points to a certain value.

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

We define a map structure with key type `int` and value type string. Then we assign it value with `{}`. Each entry is defined according `<key>: <value>` and separated by comma. So how do we read a value?

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

Imagine now that we have these two entries and you want the value given that you have then entry 555404, how would we do that? We use the square brackets like so `[]`:

```go
phonebook[555404] // "Sheriff of Nottingham"
```

### Check for existing entry

So you learned that `phonebook[555404]` gives you a value back. What if it doesn't exist? What happens if you give it a key that's not stored in the map is that you get nothing back:

```go
phonebook[888] //
```

There's a better way to check this, because accessing an entry with a key actually returns two values, the value, and a boolean. The boolean indicates if this key exist in the map. See this code:

```go
_, exist phonebook[888]
fmt.Println(exist) // false
```

You can even use this in an if statement:

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

## Assignment

TODO

## Solution

TODO

## Challenge

TODO

## Post-Lecture Quiz

TODO
