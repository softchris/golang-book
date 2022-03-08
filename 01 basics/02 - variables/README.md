# Go from the beginning - using variables

> In this second part, we cover the usage of variables in Go, how to create them, give them different types and values.

With variables, we can remember values and later refer to them via named references. using variables will make our code easier to read.

## Declare variables

In Go, there are many ways to declare variables:

- **Define a name and type**. Here, you declare a variable with the keyword `var`, give it a name and lastly a type `string`. Below is an example:

    ```golang
    var firstName string
    ```

- **Define a group** of variables. It's possible to define a grorup of variables. Using this way of declaring means you only type the `var` keyword once. The group is defined by the use of parenthesis `()`:

   ```golang
   var (
     firstName = "Chris"
     age = 20
   )
   ```

    Note how each variable is on a new row.

- **Define and assign a value**. Within functions, you can use the `:=` operator, it declares and assigns at the same time. Below code shows the creation of the `firstName` variable. The data type is inferred to be a string:

   ```go
   firstName := "Chris"
   ```

## Assign variables

To assign a new value to a variable, it needs to exist first. You use the assignment operator, `=`. Here's an example:

```go
firstName = "Mike"
```

## Data types

There are many data types you can use with Go. They are divided into different categories:

- **Basic types**. In this category, we find types like integers, floats (numbers with decimals) and other types like booleans (for true/false), strings (for text) and more.
- **Composite types**. We will talk about composite types in a separate article, but they are more complex and examples of composite types are arrays, structs and interfaces.

### Declare a variable with a type

There's two ways you can declare a variable and give it a type:

- **explicitly**, by specifying its type, for example:

   ```go
   var name string
   ```

- **implicitly**, by assigning it a value and have it be inferred:

   ```go
   name := "chris"
   ```

   In the preceding code, the data type is inferred by the value you give it. In this case, the data type becomes `string` based on the value "chris".

## Exercise - define some variables and print them out

In this exercise, you will define some variable you might need for the card game Texas Holdem.

1. Create a file *main.go* and give it the following content:

    ```go
    package main
    
    import "fmt"
    
    var (
      players = 3
      replay = false
      namePlayerOne = "chris" 
      PI = 3.14
    )
    
    func main () {
      fmt.Println(players)
      fmt.Println(replay)
      fmt.Println(namePlayerOne)
      fmt.Println(PI)
    }
    ```

1. Run `go run main.go` in the terminal:

   ```go
   go run main.go
   ``` 

   You should see the following output:

   ```output
    3
    false
    chris
    3.14
   ```

## String interpolation

Sometimes, you want to be able to write things to the screen and mix different data types doing so. For example, you might want to write, "Customer: Adam has 20$ on his bank account".

Lets say then that this information is represented by the these two variables:

```go
var (
  customerName = "Adam"
  accountBalance = 20
)
```

How can you print out the text above? For this purpose, you can use the `Printf()` function that takes formatters. The idea is that a formatter is an instruction to what a certain type is. By providing this information to `Printf()`, it's able to print the type correctly.

Here's how you can print the example string from before:

```go
fmt.Printf("Customer %s has %d$ on their bank account", customerName, accountBalance)
```

Above, the `%s` represents a string and `%d` represents a number. By using these formatters as placeholders, the variables are correctly implemented and the output becomes:

```output
Customer Adam has 20$ on their bank account
```

## Summary

In this exercise, you've learned how to declare your first variables. You've also learned that there's different types for different types of data. Finally you've learned too interpolate variables of different types and have that be printed to the screen.

### Learn more

<https://docs.microsoft.com/en-us/learn/modules/go-variables-functions-packages/2-data-types>
