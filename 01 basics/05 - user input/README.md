# Reading user input

It's an important thing to be able to read user input from the console. How to do that is by using the built-in `fmt`.

The `fmt` library has a built-in `Scan()` method.

## Reading one input

Here's how to use it:

```go
package main

import "fmt"

func main() {
    var response string
    fmt.Scan(&response)
    fmt.Println("User typed", response)
}
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