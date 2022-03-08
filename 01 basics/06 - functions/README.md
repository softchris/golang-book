# Functions

So far, you've seen the function `main`, define like so:

```go
func main(){
}
```

There's only one such function, it's called an entry point and represents the start of the program. You can however define other functions.

To create a function, you need:

- `func`, the keyword `func`.
- **parameters**, 0 to many parameters
- **a function body**, i.e statements that says what the function does
- **a return construct**, if the function returns something

## A first function

Let's show a function:

```go
package main

import "fmt"

func log() {
  fmt.Println("message")
}

func main() {
  log()
}
```

This program would output "message". To make `log()` more flexible, lets add a parameter.

### Adding a parameter

```go
func log(message string) {
  fmt.Println(message)
}

// to use
log("hi")
log("there")
```

Note how the `log()` function takes the parameter `message` that is of type string. Our code is more flexible.

### Adding a return type

To add a return type, we add that after parenthesis in form of a type. Here's an example:

```go
add(firstNumber int, secondNumber int) int {
 ...
}
```

Because we've added a return type of `int`, our function must return something. A way to return a value is by using the keyword `return`, like so:

```go
add(firstNumber int, secondNumber int) int {
  return firstNumber + secondNumber 
}
```

### Named return

We can also name the return parameter like so:

```go
add(firstNumber int, secondNumber int) (sum int) {
  sum = firstNumber + secondNumber 
  return
}
```

- Note how `sum` is part of function prototype declaration `(sum int)` and then assigned a value `sum = firstNumber + secondNumber`.

- There's also a `return` on its own row, this code will compile as there's a notion of a return variable.

### Multiple returns

It's possible to return more than one value.

Just like you returned a named parameter via `(sum int)`, you can comma separate like so `(sum int, product int)`. When return the values, you can type like so:

```go
sum = first + second
product = first * second
return
```

Both `sum` and `product` are assigned values and you have a closing `return`. 

Putting it altogether you get a function that looks like so:

```go
func calc(first int, second int) (sum int, product int) {
 sum = first + second
 product = first * second
 return
}
```

To call the function, you type like so:

```go
sum, product = calc(1, 2)
fmt.Println(sum)
fmt.Println(product)
```

## Command line args

A way to get input to your program is via the command line. The `os` library allows us to read command line arguments. Here's how to use it:

```go
package main
import (
  "os"
  "fmt"
)

func main() {
  os.Args[1] 
}
```

`os.Args` represents an array of the command line arguments. Let's say the above program is called like so:

```bash
main 1 2
```

Then the arguments will be placed like the following in `os.Args`:

```go
os.Args[0] // program name
os.Args[1] // "1"
os.Args[1] // "2"
```
