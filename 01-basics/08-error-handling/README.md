# Handling errors

In this chapter, we will cover error handling.

## Introduction

This chapter will cover:

- Causing and handling errors with `panic` and `recover`.
- Use the error pattern to use a more idiomatic approach to managing errors.

## Errors

Our apps aren't perfect, they will throw errors. Either they throw errors because of code we wrote or because of code written in a package we are using. Sometimes, that's even what the code is supposed to do when we feed it bad input.

Regardless, our users expect us to handle these errors. Also, for our own sake, we need to log these errors in such a way that we can fix those errors if they are unexpected.

There are two major ways to handle errors in Go:

- **panic/recover**. `panic` and `recover` are language constructs. If you know another programming language, they function like throw and catch. What does it mean though? It means there's an error, with an error message and stack trace that we can catch and recover from or let it crash the program.
- **error pattern**. This is referred to as the idiomatic Go way of doing things. The idea is to return a value and an error object from a function call. If an error occurred, it contains an error object or `nil`, if no error.

## Crash the program with `panic()`

Let's take this function `Divide()`:

```go
func Divide(nominator int, divider int) float32 {
  if divider == 0 {
    panic("can't divide by 0")
  }
  return float32(nominator) / float32(divider)
}
```

It has an `if` check. If `divider` is `0` then it calls `panic()`. So what happens then? You see something like:

```output
panic: can't divide by 0

goroutine 1 [running]:
main.Divide(...)
        /<path>/panic.go:20
main.main()
        /<path>/panic.go:33 +0x96
exit status 2
```

At the top is the error string you sent to `panic()`, the string "can't divide by 0". Then you have the stack trace, entries that indicate where the error started. Read it from the bottom, the error started on line 33, then you have line 20, which is this row:

```go
panic("can't divide by 0")
```

Ok, so we have a way to protect ourselves from input values that we don't want. But crashing the program, is that necessary? In some cases, it is, in some cases, it isn't. For the latter cases, we have `recover()`.

### Capture the error with `recover()`

Using `recover()` is about capturing an error so our program can continue. We need to learn about a concept before proceeding. That concept is `defer`. `defer` is a language construct that delays the execution of a function until the nearby function returns. Here's an example:

```go
defer fmt.Println("second")
fmt.Println("first")
```

Running this program, you should see:

```go
first
second
```

> See `defer` as the last thing that happens.

How is this useful in the context of capturing an error? Capturing errors, if you want to capture it, is something you want to do as the last thing that happens. Take our `Divide()` function again:

```go
func Divide(nominator int, divider int) float32 {
  defer errorHandler()
  if divider == 0 {
    panic("can't divide by 0")
  }
  return float32(nominator) / float32(divider)
}
```

Note how it now has a line in it that says `defer errorHandler()`. It will be run the last thing that happens. Depending on the value of `divider`, it will either call `panic()` or call the `return` statement as the last thing.

Ok, so what does `errorHandler()` look like?

```go
func errorHandler() {
  if r := recover(); r != nil {
    fmt.Println("Recovered ", r)
  }
}
```

In `errorHandler()`, we invoke `recover()` and assign it to variable `r` and then we test it for `nil`. If it's NOT `nil` then we have an error and we print it out. If it is `nil` then the user notices nothing.

## Improve the error handling

There are steps we can take to improve the error handling. So far, we're printing back an error message. Imagine if someone read this error from a log file, would they be able to understand where things went wrong and fix the input data or the code itself?

There are a couple of things we can do:

- **Inspect the error**. Our error didn't only come with an error message, there's a stack trace as well.
- **Logging**. If we want someone to work with these errors, we should look at logging them to a file.

### Inspect the error with `runtime/debug`

There's a library `runtime/debug`. With this library, you can find out more about an error when it's thrown, information like stack trace, where the error originated. There's a function `Stack()` that produces the stack trace. Here's how to use it:

```go
debug.Stack()
```

### Log the error with `log`

While `runtime/debug` can produce a stack trace, what would be useful is logging all this error information to a file. To log to a file, use the `os` and `log` package like so:

```go
f, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
 if err != nil {
  log.Println(err)
 }
 log.SetOutput(f)
```

## Use the error pattern with the `errors` package

Other languages tend to use Exceptions to signal that something is wrong.

Go has a different and idiomatic approach. It wants you to create errors as return values to a function, next to the actual value being returned. You are then expected to inspect a function and see if it returns an error.

There's an `errors` package that can help us with the above approach.

### Define an error

To define an error, we call the `New()` function with a string describing the error, here's an example:

```go
var NoTooSmall = errors.New("the number is too small")
```

Next, let's look at how to add the error to a function.

### Return an error

Let's start with function that uses a `panic()` as error handling:

```go
func ReturnPositive(no int) int {
  if no > 0 {
    return no
  } else {
    panic("No too small")
  }
}
```

We can improve this function, by ensuring it always returns the result and an error, like so:

```go
func ReturnPositive(no int) (int, error) {
  if no > 0 {
    return no, nil
  } else {
    return 0, NoTooSmall
  }
}
```

Note in the `if` clause that it returns `no` and `nil` when everything is fine. For the `else`, it returns a bogus value and our error `NoTooSmall`.

### Inspect the result

Let's see how we would call the `ReturnPositive()` function and use this new pattern we established:

```go
no, err := ReturnPositive(-2)
if err != nil {
  fmt.Println("error: ", err)
} else {
  fmt.Println("value:", no)
}
```

What you are seeing above is how we use an `if` clause to check for errors, if so, print out. On the `else`, we have our actual value.

## Assignment I - add error handling

In this exercise, we'll add error handling to our program.

1. Create a file *panic.go* and give it the following content:

   ```go
   package main

   import (
     "fmt"
   )
    
    // func errorHandler() {
    // if r := recover(); r != nil {
    //  fmt.Println("Recovered ", r)
    // }
    // }
    
    func Divide(nominator int, divider int) float32 {
     // defer errorHandler()
     if divider == 0 {
      panic("can't divide by 0")
     }
     return float32(nominator) / float32(divider)
    }
    
    func main() {
     no := Divide(10, 0)
     fmt.Println(no)
     no = Divide(10, 1)
     fmt.Println(no)
    }
   ```

1. Run the program `go run panic.go`

   ```bash
   go run panic.go
   ```

   You should see output similar to:

   ```output
   panic: can't divide by 0

    goroutine 1 [running]:
    main.Divide(...)
            /<path>/panic.go:20
    main.main()
            /path/panic.go:33 +0x96
    exit status 2
   ```

   Note how these two statements was never run as the program crashed:

   ```go
   no = Divide(10, 1)
   fmt.Println(no)
   ```

1. Uncomment the commented out part and run the app again.

   You should now see the following output:

   ```output
   Recovered  can't divide by 0
   0
   10
   ```

Congrats, you've managed to implement error handling with `panic()` and `recover()`.

## Solution I

```go
package main

import (
   "fmt"
)
   
func errorHandler() {
   if r := recover(); r != nil {
      fmt.Println("Recovered ", r)
   }
}
   
func Divide(nominator int, divider int) float32 {
  defer errorHandler()
  if divider == 0 {
    panic("can't divide by 0")
  }
  return float32(nominator) / float32(divider)
}
   
func main() {
  no := Divide(10, 0)
  fmt.Println(no)
  no = Divide(10, 1)
  fmt.Println(no)
}
```

## Assignment II - improve error logging

Let's improve our *panic.go* file by adding error logging:

1. Locate the `errorHandler()` and change it to the following:

   ```go
   func errorHandler() {
     if r := recover(); r != nil {
       log.Println(r, string(debug.Stack()))
     }
   }
   ```

1. Ensure that *panic.go* looks like so:

   ```go
    package main

    import (
     "fmt"
     "log"
     "os"
     "runtime/debug"
    )
    
    func errorHandler() {
     if r := recover(); r != nil {
      log.Println(r, string(debug.Stack()))
     }
    }
    
    func Divide(nominator int, divider int) float32 {
     defer errorHandler()
     if divider == 0 {
      panic("can't divide by 0")
     }
     return float32(nominator) / float32(divider)
    }
    
    func main() {
     f, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
     if err != nil {
      log.Println(err)
     }
     log.SetOutput(f)
    
     log.Println("starting program")
     no := Divide(10, 0)
     fmt.Println(no)
    
     no = Divide(10, 1)
     fmt.Println(no)
     f.Close()
    }
   ```

1. Run this program:

   ```go
   go run panic.go
   ```

   You should see this output:

   ```output
   0
   10
   ```

   It was able to run all statements without being affected by the error that was thrown.

1. Inspect the *logs* file that was just created, it should have content like the below:

   ```output
     2022/03/11 15:03:59 starting program

    2022/03/11 15:03:59 can't divide by 0 goroutine 1 [running]:
    runtime/debug.Stack(0xc000111d30, 0x10b1b40, 0x10eae78)
     /usr/local/Cellar/go/1.16/libexec/src/runtime/debug/stack.go:24 +0x9f
    main.errorHandler()
     /<path>/panic.go:14 +0x5b
    panic(0x10b1b40, 0x10eae78)
     /usr/local/Cellar/go/1.16/libexec/src/runtime/panic.go:965 +0x1b9
    main.Divide(0xa, 0x0, 0x0)
     /<path>/panic.go:21 +0xa5
    main.main()
     /<path>/panic.go:34 +0x115
   ```

## Solution II

```go
package main

import (
  "fmt"
  "log"
  "os"
  "runtime/debug"
)

func errorHandler() {
  if r := recover(); r != nil {
    log.Println(r, string(debug.Stack()))
  }
}

func Divide(nominator int, divider int) float32 {
  defer errorHandler()
  if divider == 0 {
    panic("can't divide by 0")
  }
  return float32(nominator) / float32(divider)
}

func main() {
  f, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
  if err != nil {
    log.Println(err)
  }
  log.SetOutput(f)

  log.Println("starting program")
  no := Divide(10, 0)
  fmt.Println(no)

  no = Divide(10, 1)
  fmt.Println(no)
  f.Close()
}
```
