# Converting between types

This chapter covers how to convert between strings and numbers.

## Introduction

This chapter will:

- Introduce uses cases where data conversion makes sense.
- Showcase how to use `strconv` library.

## Why convert between types

There are different data types and a need to convert between them. For example, we often need to convert between text and numbers for presentational and other reasons. We also need to convert between numbers and decimals without losing information in the process.

The main package for dealing with conversions in Go is `strconv`.

## Use case - command-line arguments

Let's show a common case where you start off with strings and you need to make it into numbers, command-line arguments. To use command-line arguments in a program, you need the `os` package.  

`os.Args` points to an array representing your command line arguments. To access a specific argument, you would use the index operator `[]` like so:

```go
arg := os.Args[1]
```

You can then start your program like so:

```bash
go run main.go 1
```

The 1 would then be stored in `arg`.

### Finding the type

What type is `arg` in our code above? There are some ways to find out:

- **IDE**, if you use for example Visual Studio Code and the Go plugin, hovering over the code, it will tell you that `os.Args` is a string array, `string[]`.
- **PrintF() and %T**. One of the easiest ways to find the type is typing like so:

   ```go
   Printf("%T", os.Args[1])
   Printf("%T", 1)
   ```

   You would get an output like so:

   ```output
   string
   int
   ```

- **Type coercion**, you could try to modify that code and coerce it to be an integer like so, now what?

    ```go
    var no int = os.Args[1]
    ```

    You get an error:

    ```output
    cannot use os.Args[1] (type string) as type int in assignment
    ```

- **Use reflection**. Another way to find the above is by using the `reflect` package like so:

    ```go
    package main
    
    import (
      "reflect"
      "fmt"
      "os"
    )
    
    func main () {
      arg := os.Args[1]
      fmt.Println(reflect.TypeOf(arg))
    }
    ```

    Now, the program will print "string" as the type.

### Addressing the problem with `strconv`

Ok, so we know what type something is, what if we need to use these command-line arguments, which are of type `string`, and feed them into let's say a calculator program?

Consider the below code, that at present WOULDN'T compile:

```go
package main

import (
  "fmt"
  "os"
)

func add(first int, second int) int {
  return first + second
}

func main() {
  add(os.Args[1], os.Args[2]) // this would NOT compile
}
```

The reason is that the values on `os.Args[1]` and `os.Args[2]` are `string` not `int`. To fix this issue, we need to use the conversion package `strconv`.

## Convert from string to int with `strconv`

To convert strings to integers, we need to use `strconv` and call the `Atoi()` (stands for Ascii to integer) function like so:

```go
package main

import (
 "fmt"
 "os"
 "strconv"
)

func add(first int, second int) int {
 return first + second
}

func main() {
 no1, _ := strconv.Atoi(os.Args[1])
 no2, _ := strconv.Atoi(os.Args[2])
 var sum = add(no1, no2) 
 fmt.Println(sum)
}
```

Note `_`, this is a *don't care* symbol. What happens when you call `Atoi()` is that it returns two things, the number and an error if it fails.

### Handling conversion error

To handle an error, we need to store it in a variable, `err` and inspect it. If it's not `nil`, then we have an error.

Here's how we could encode that behaviour below:

```go
package main

import (
 "fmt"
 "os"
 "strconv"
)

func main() {
 no, err := strconv.Atoi(os.Args[1])
 fmt.Println(no)
 if err != nil {
  fmt.Println(err)
  fmt.Println("Couldn't convert: " + os.Args[1])

 } else {
  fmt.Println(no)
 }
}
```

Try to compile the above program and run it like so:

```bash
main 1 # 1
main hi # trconv.Atoi: parsing "hi": invalid syntax, Couldn't convert: hi
```

## Parse string to int

There's another way to convert a string to an int. That's by using the `ParseInt()` method. It does more than converting though, it does two things in fact:

- **base**, you can select according to what base to interpret the number as.
- **size**, bit size, from 0 to 64.

The syntax for the method looks like so:

```go
ParseInt(<s string>, <base int>, <bit int>) (int64, error)
```

Here's some examples:

```go
package main

import (
 "fmt"
 "reflect"
 "strconv"
)

func main() {
 var no int = 100
 fmt.Println(reflect.TypeOf(no))

 var intStr string = "100"
 fourBaseEightBitInt, _ := strconv.ParseInt(intStr, 4, 8)    // becomes no 16 and int64
 tenBaseSixteenBitInt, _ := strconv.ParseInt(intStr, 10, 16) // no 100,  and int64
 fmt.Println(reflect.TypeOf(fourBaseEightBitInt))
 fmt.Println(reflect.TypeOf(tenBaseSixteenBitInt))
}
```

## Integer to string

You might be dealing with the opposite; you have an integer, and you want it to be a string. In this case, you can use the `Itoa()` function, integer to ascii. Here's an example:

```go
var noOfPlayers = 8
str := strconv.Itoa(noOfPlayers)

```

## Additional parsing

The `strconv` library is what you want if you start with a string, and you want to convert to and from another format. Learn more about [strconv library here](https://pkg.go.dev/strconv)

## Assignment

Create an app that adds two numbers together. The values should come from the command line. Here's how the program should run:

```bash
go run main.go 2 4
6
```

## Solution

```go
package main

import (
 "fmt"
 "os"
 "strconv"
)

func add(no int, secondNumber int) int {
 return no + secondNumber
}

func main() {
 no1, _ := strconv.Atoi(os.Args[1])
 no2, _ := strconv.Atoi(os.Args[2])
 var sum = add(no1, no2)
 fmt.Println(sum)
}
```

## 🚀 Challenge

What happens if run the program like so?

```bash
go run main.go one two
```

Handle any conversion error and call `panic()` if there's a conversion error.
