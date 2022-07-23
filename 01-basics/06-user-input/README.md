# Reading user input

You will learn how to read user input, both a simpler technique and a more advanced one using formatters.

## Introduction

This chapter will cover:

- The `Scan()` method to read user input.
- Reading one input.
- Reading multiple inputs.
- Working with formattters.

## User input

It's an important thing to be able to read user input from the console. It gives the user a chance to interact with the program. Things to consider are how you are asking the user to input, is it one word, several inputs. Will the user separate input by space or newline? Regardless of what approach you go with, try to communicate the chosen approach to the user.

## Managing user input with `fmt`

So far, you've seen how the `fmt` package can be used to print to the console. It can also be used to read user input.

The `fmt` library has a built-in `Scan()` method that we will use to capture the user input.

## Reading one input

You might start out wanting to read one input from the user. That's what the `Scan()` method does by default.

Here's some code showing how to collect user input:

```go
package main

import "fmt"

func main() {
    var response string
    fmt.Scan(&response)
    fmt.Println("User typed: ", response)
}
```

Note how you send in the string `response` as a reference, using the `&` operator. It's done this way as the `Scan()` method will modify the variable you send in.

When you run this code, you will see the below output:

```output
hello
User typed: hello
```

## Reading multiple inputs

You can provide several arguments to the `Scan()` method. By providing several arguments, you can collect more than one user input and separate each input, in the `Scan()` function, with a comma. Here's how to apply this technique:

```go
var a1, a2 string
// multiple input
fmt.Scan(&a1, &a2)

// formatted string to print out the user input
str := fmt.Sprintf("a1: %s a2: %s", a1, a2)
```

Note how `a1` and `a2` is sent into `Scan()` and they are comma-separated. So how will those codes run?

When you run this code, there are two ways for the user to input. The user can either separate the values by space like so:

```output
hi you
a1: hi a2: you
```

or by newline:

```output
hi 
you
a1: hi a2: you
```

## `Scanf()`, scan the input with formatters

So far, you've seen how you can collect input with spaces and endlines as separators. You can also be specific in how you collect input. Imagine that the user wants to type in an invoice number like so "INV200" or "INV114" and what you are interested in is the number part, or you are interested in the prefix as well?

The answer to solving this is in the `Scanf()` function. It takes formatters. With formatters, you're able to pick the part of the user input you are interested in and place that into the variable you want.

In the above-mentioned case, you can construct code looking like so:

```go
var prefix string
var no int
// inv110
fmt.Scanf("%3s%d", &prefix, &no)
fmt.Printf("prefix: %s, invoice no: %d", prefix, no)
```

The interesting part lies in the first argument of `Scanf()`, namely `3s%d`. What this means is, take the first 3 characters, `%3s` and interpret them as a string. Then interpret and place the remaining as a number, `%d`.

Running this program, you will get an output like so:

```output
inv200
prefix: inv, invoice no: 200
```

"inv" is placed in `prefix` and 200 in `no` variable.

## Learn more

To learn more about this area, check out this link <https://pkg.go.dev/fmt#Scanf>

## Assignment - read user input

In this assignment, you will read user input for a card game. The objective is to capture the names of the players.

1. Create a file *main.go* with the following content:

   ```go
   package main

   import "fmt"

   func main() {

   }
   ```

1. Add the following code to the `main()` method:

   ```go
   fmt.Println("Enter player names (separated by space or newline):")
   var player1, player2 string
   fmt.Scan(&player1, &player2)
   fmt.Println("Players are: ", player1, player2)
   ```

1. Run the program using `go run main.go`:

   ```bash
   go run main.go
   ```

   You should see the following output:

   ```output
   Enter player names (separated by space or newline):
   Alice Bob
   Players are:  Alice Bob
   ```

## 🚀 Challenge

Try modifying the program so it takes several players first and then ensures all players get a name. Here's an example output of such a program:

```output
Enter number of players: 
3
Enter Player 1 name: 
Alice
Enter Player 2 name:
Bob
Enter Player 3 name:
Jean
Players are: Alice Bob Jean
```  

## Solution

```go
package main

import "fmt"

func main() {
 fmt.Println("Enter player names (separated by space):")
 var player1, player2 string
 fmt.Scan(&player1, &player2)
 fmt.Println("Players are: ", player1, player2)
}
```
