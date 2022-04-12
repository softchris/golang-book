# Your first project

In this chapter we will cover how to create your first project in Go.

##  Pre-Lecture Quiz

TODO

## Introduction

This chapter will cover:

- Creating a project
- Organize your files

## Module use cases

There are two interesting use cases with modules:

- **Consuming a module**, you will most likely use a combination oof core modules and external 3rd party modules
- **Creating a module**, in some case you will create code that you or someone else will be able to use. For this scenario, you can create a module and upload it to GitHub.

## Consume internal files

You want to split up your app in many different files. Let's say you have the following files:

```output
/app
  main.go
  /helper
    helper.go 
```

What you are saying above is that your program consist of many files and that you want code in *main.go* to use code from *helper.go* for example.

To handle such a case, you need the following:

- **a project**. By creating a project, you create a top level reference that you can use in the `import` directive.
- **an import** that points to the project root name as well as the path to the module you want to important.

You can use `go mod init`, this will initialize your project.

## Creating a project

To create a project, you run `go mod init` and a name for a project, for example "my-project":

   ```bash
   go mod init my-project
   ```

   You end up with a *go.mod* file looking something like so:

   ```go
   module my-project

   go 1.16
   ```

   The *go.mod* file tells you the name of your project and the currently used version of Go. It can contain other things as well like libraries you are dependent on.

## The import statement

Imagine now we have this file structure in our project:

```output
/app
  main.go
  /helper
    helper.go 
```

with *helper.go* looking like so:

```go
package helper
    
import "fmt"

func Help() {
  fmt.Println("This is a helper function")
}
```

to use the public `Helper()` function from *main.go*, we need to import it.

In *main.go* we need an import statement like so:

```go
import (
  "my-project/helper"
)
```

We are now able to invoke the `Help()` function from *main.go* like so:

```go
helper.Help()
```

## Assignment - create a project

In this assignment

1. Create the module like so:

    ```go
    go mod init
    ```

1. create the **helper** directory and *helper.go* file and give it the following content:

    ```go
    // helper.go
    
    package helper
    
    import "fmt"
    
    func Help() {
     fmt.Println("This is a helper function")
    }
    ```

1. Create the *main.go* file and give it the following content:

   ```go
   package main

   import (
     "log-tester/helper"
   )
    
   func main() {
     helper.Help()
   }
   ```

   Note this import `"log-tester/helper"`, it ensures the `helper` package is in scope.

1. Compile and run

   ```bash
   go run main.go
   ```

## Solution

helper/helper.go

```go
package helper

import "fmt"

func Help() {
 fmt.Println("help")
}
```

main.go

```go
package main

import "my-project/helper"

func main() {
 helper.Help()
}
```
