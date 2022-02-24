# Your first program

## Concepts

- `package main`, the entry point needs to have this instruction
- `import "fmt"`, fmt is standard package for input and ouput
- `func main`, entry point function, where your program starts

### Run your app

To run your app, type `go run <file>.go`, for example:

```bash
go run main.go
```

### Build your app

To produce an executable, run `go build <file>.go`, for example:

```bash
go build main.go
```

It produces an executable, on MacOS and Linux that's a file with -X as permission, on Windows, it's a .exe file.

## Code

Here's what a first program can look like:

```golang
package main

import "fmt"

func main() {
 fmt.Println("hello")
}
```
