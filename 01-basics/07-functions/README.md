# Using functions

In this chapter, we will discuss how you can define and use functions. Functions are great when you have the same type of code used in many places. By using functions, you thereby reduce repetition.

## Introduction

This chapter will cover:

- What a function is and why use it.
- How to define and call a function.
- Returning multiple values.

## Why functions

As soon as you have a set of statements you repeat in many places, it's a good use case for creating a function. Typical things you put in functions are logging to a file, performing a calculation or talking to a data source.

## Your first function `main()`

So far, you've seen the function `main()`, define like so:

```go
func main(){
}
```

There's only one such function, it's called an entry point and represents the start of the program. You can however define other functions.

## The anatomy of a function

A function consists of various parts. By incorporating all these parts, you ensure you have a reusable piece of code you can use in many places.

Here are the parts you need to care about:

- `func`, the keyword `func`.
- **parameters**, 0 to many parameters
- **a function body**, i.e. statements that say what the function does.
- **a return construct**, if the function returns something.

Here's an example:

```go
func add(first int, second int) int {
  return first + second
}
```

In the preceding code, the function is named `add()`. It has the parameters `first` and `second`. The function body, what the function does, consists of this code:

```go
return first + second
```

## Adding a return type

To add a return type, we add that after the function parenthesis in form of a type. Here's an example:

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

## Multiple returns

It's possible to return more than one value.

Just like you returned a named parameter via `(sum int)`, you can comma separate like so `(sum int, product int)`. When returning multiple values, you can type like so:

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
sum, product := calc(1, 2)
fmt.Println(sum)
fmt.Println(product)
```

Note how you assign the two returned values to variables `sum` and `product`.

## Assignment - adding a function to a program

1. Create a file *main.go* and give it the following content:

    ```go
    package main
    
    import "fmt"
    
    func main() {
    
    }
    ```

1. Add a function `log()`, that we can use to print messages.

   Added to the program, your code should now look like so:

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

At this point, the `log()` function isn't very flexible, it prints "message" every time it's invoked.

To make the `log()` function more flexible, lets add a parameter.

### Adding a parameter

A parameter needs a data type, in this case, we will make it of type `string`.

1. Add the parameter within the parenthesis `()`, like so:

    ```go
    func log(message string) {
      fmt.Println(message)
    }
    
    // to use
    log("hi")
    log("there")
    ```

Note how the `log()` function takes the parameter `message` that is of type string. Our code is more flexible.

## Solution

```go
package main
    
import "fmt"

func log(message string) {
  fmt.Println(message)
}

func main() {
  log("hi")
  log("there")
}
```
