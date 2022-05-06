# Working with strings

There are many reasons why we need to work with strings in different ways. Below are some situations that you are likely to encounter and that the `strings` library has a solution for:

- **user input**, or other types of storage may contain special characters that you need to cater for.
- **inspection**, does the string contain what we need for our business logic?
- **parsing**, splitting the file until we get what we need, could be things like a number or date for example.
- **presentation**, sometimes we need to present the text in certain way in a UI for example, to highlight said information.

## Handling special characters

Lets say we read user input and we want to interpret what we get as a number. To make our program robust, we're ok with the user typing spaces or newline characters. The following input should be allowed:

```text
114
   114
114\n
```

There's three methods of interest to handle such a case for us, namely `Trim()`, `TrimLeft()` and `TrimRight()`.

- `Trim()`, what it does is remove whitespace characters from both left and right side.
- `TrimLeft()`. It removes whitespace from the left side only, and if you specify special characters as well.
- `TrimRight()`. It removes whitespace from the right side only, and if you specify special characters as well.

All these functions above have as their second parameter a so called *cutset* parameter, where you specify what character you want to get rid of. You can for example specify to remove space, newline and tab characters like so:

```go
" \n\t"
```

Here's some code that shows all three methods in use:

```go
fmt.Printf("%s , string length %d \n", s, len(s))
 res := strings.Trim(s, " ")
 fmt.Printf("%s , string length %d \n", res, len(res))

 s2 := "   114  "
 fmt.Printf("%s , string length %d \n", s2, len(s2))
 res = strings.TrimLeft(s2, " ")
 fmt.Printf("%s , string length %d \n", res, len(res))

 s3 := "   114  "
 fmt.Printf("%s , string length %d \n", s3, len(s3))
 res = strings.TrimRight(s3, " ")
 fmt.Printf("%s , string length %d \n", res, len(res))
```

The above string has three whitespaces to the left and two to the right, giving it a total length of 8.

The output of the above code is:

```output
  114   , string length 8 
114 , string length 3 
   114   , string length 8 
114   , string length 5 
   114   , string length 8 
   114 , string length 6 
```

Lets break down the output per row.

For this output, using `Trim()`, the spaces are removed on both sides and we end up with something looking left aligned:

```output
"114"
```

The next output, using `TrimLeft()`, shows how the right spaces are still there:

```output
"114   "
```

Our final row, using `TrimRight()`, shows a right alignment and how the spaces on the left side still remains:

```output
"   114"
```

## Inspect with `Contains()`

Imagine you want to inspect a string to verify whether it contains a certain substring.

For that, you can use the `Contains()` function. Its syntax looks like so:

```go
strings.Contains(stringSource, pattern)
```

You can then for example use it to process a list from a point of sale system, and for each item check if it contains a certain prefix:

```go
rows := []string{"order: 5", "order: 10", "order: 5", "separator"}


for item :=  range rows {
  if strings.Contains("order") {
    // process order
  }
  // ignore
}
```

## Parsing with `Split()`

Lets continue with processing rows from our point of sale system. This time, we will be looking at a specific item and extract the information we need. For this, we will use the `Split()` function:

```go
rows := []string{"order: 5", "order: 10", "order: 5", "separator"}

for item :=  range rows {
  if strings.Contains("order") {
     tokens := strings.Split(item, ":") // [ "order", " 5"]
     value := strings.Trim(tokens[1])
     fmt.Println(value)
  }
  // ignore
}
```

By using this code:

```go
strings.Split(item, ":")
```

on this string "order: 5", we end up with an array `["order", "5"]` and `strings.Trim(tokens[1])` would then refer to 5.

> TIP: If we need to treat the 5 above as a number, as part of calculation, we would need to convert it to a number first

## Presentation

Say you have customer management system and there's a lot of data to present. To give importance to certain data, we can use functions to highlight their visual appearance.

Take the following multiline customer string:

```text
Jean Normand
123 Way
Washington
```

If you use `ToUpper()` on city you get a result like so:

```output
Jean Normand
123 Way
WASHINGTON
```

With `ToLower()` you ensure all characters are formatted as lowercase.

## Assignment

Write a program that given a struct containing, name, address and city ensures that the name is lowercase and the address is uppercase.

## Solution

```go
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
 person := Person{Name: "jean Normand", Address: "123 Way", City: "Washington"}

 fmt.Println(strings.ToUpper(person.Name))
 fmt.Println(person.Address)
 fmt.Println(strings.ToUpper(person.City))
}
```
