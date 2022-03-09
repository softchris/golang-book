# Reading user input

It's an important thing to be able to read user input from the console. How to do that is by using the built-in `fmt`.

The `fmt` library has a built-in `Scan()` method.

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

Not how you send in the string `response` as a reference, using the `&` operator. It's done this way as the `Scan()` method will modify the variable you send in. 

When you run this code, you will see the below output:

```output
hello # this is what we typed
User typed: hello
```

## Reading multiple inputs

You can provide several arguments to the `Scan()` method. By providing several arguments, you can collect more than one user input and separate each input with a comma. Here's how to apply this technique:

```go
var a1, a2 string
// multiple input
fmt.Scan(&a1, &a2)

// formatted string
str := fmt.Sprintf("a1: %s a2: %s", a1, a2)
```

Note how `a1` and `a2` is sent into `Scan()` and they are comma separated, so how will those code run?

When you run this code, there's two ways for the user to input. The user can either separate the values by space like so:

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

## Scanf, scan the input with formatters

So far, you've seen how you can collect input with spaces and endlines as separators. You can also be very specific in how you collect input. Imagine that the user wants to type in an invoice number like so "INV200" or "INV114" and what you are interested in is the number part, or maybe you are interested in the prefix as well? The answer to solving this is in the `Scanf()`  function. It takes formatters. With formatters, you're able to pick the part of the user input you are interested in, and place that into the variable you want. 

In the above mentioned case, you can construct code looking like so:

```go
var prefix string
var no int
// in110
fmt.Scanf("%3s%d", &prefix, &no)
fmt.Printf("prefix: %s, invoice no: %d", prefix, no)
```

The interesting part lies in the first argument of `Scanf()`, namely `53s%d`. What the means is, take the first 3 characters, `%3s` and interpret as a string. Then interpret and place the remaining as a number, `%d`. 

Running this program you will get an output like so:

```output
inv200
prefix: inv, invoice no: 200
```

"inv" is placed in `prefix` and 200 in `no` variable.

## Learn more

To learn more about this are, check out this link <https://pkg.go.dev/fmt#Scanf>
