# Boolean logic

> This article will cover working with boolean logic. You will learn how to work with boolean data, `if`, `else if` and `else` constructs.

Using boolean logic in your program is about creating different execution paths through your code?

> What does that mean?

It means there's more than one way that your program can run depending on what data you feed it.

> Ok, can you show me?

Sure, consider this code:

```go
printMessage := true
 if printMessage {
  fmt.Println("Message")
 }
```

If `printMessage` is `true`, the string "Message" will print. If the value is `false`, nothing will print.

> Ok, I think I get it.

## The `if` construct

You've seen an example already about code that runs or don't run depending on a value. The `if` construct is what makes that possible. An `if` take a boolean expression like so:

```go
if true {
  // statements here will always run
}
```

### Using a boolean variable

When you use a boolean value as part of your boolean expression, it needs to be evaluated. Here's code showing just that:

```go
accountBalance = 100
accountCredit = 200
if accountBalance + accountCredit > 0 {
  fmt.Println("You have money to spend")
}
```

The program above does the job, meaning it correctly evaluates whether you have money to spend. However, you might want to print something out if the condition is not met, for that you have `else`.

### introducing `else`

You would like to improve the preceding code. The `else` clause is run when `if` is evaluated to false. Here's how you can add it to the program:

```go
accountBalance = 100
accountCredit = 200
if accountBalance + accountCredit > 0 {
  fmt.Println("You have money to spend")
} else {
  fmt.Println("No money left, please add more funds")
}
```

## Using `else if`

`if` and `else` takes you far. Sometimes, it's not enough. You might need to grade a course in different levels depending on the points achieved on the exam. For this situation, you need an `else if` construct, a construct that will be evaluated if the `if` constructs evaluates to false. It differs from `else` in that it also takes an expression. Here's an example where it's used:

```go
if testScore >= testScoreGrade5 {
  fmt.Println("Top mark")
 } else if testScore >= testScoreGrade4 {
  fmt.Println("Pass with distinction")
 } else if testScore >= testScoreGrade3 {
  fmt.Println("Pass with distinction")
 } else {
  fmt.Println("Failed")
 }
```

## Multiple expressions

Your expression can examine more than one variable or condition. There are boolean operators you can use to help you. Here are some operators you are likely to encounter:

- `&&`, evaluates to true if values on the left and right side are both true. Here's an example of this operator in use:

    ```go
    hasGas := true
    hasKeyInIgnition := true
    if hasGas && hasKeyInIgnition {
      fmt.Println("Can drive car")
     }
    ```

    In the preceding code, the expression will evalute to true as both `hasGas` and `hasKeyInIgnition` is true.

- `||` , evaluates to true if either left or right value is true. Here's an example of this operator in use:

  ```go
  hasBurger := true
  hasSandwich := false

  if hasBurger || hasSandwich {
    fmt.Println("Can eat")
  }
  ```

  In the preceding code, `hasBurger` is true and that's enough for this expression to become true.

- `!`, also know as NOT, it will negate the expression. Here's an example:

   ```go
   hasSandwich := false

   if !hasSandwich {
    mt.Println("No sandwiches, then I will starve, I only eat sandwiches")
   }
   ```

   Above, the expression will evaluate to true, thanks to the negation with `!`.

## Exercise - create a program that tests your boolean logic

1. Create a file *main.go* and give it the following content:

   ```go
    package main

    import "fmt"
    
    func main() {
     testScoreGrade5 := 80
     testScoreGrade4 := 60
     testScoreGrade3 := 50
     testScore := 49
    
     hasGas := true
     hasKeyInIgnition := true
    
     hasBurger := true
     hasSandwich := false
    
     printMessage := true
     if printMessage {
      fmt.Println("Message")
     }
    
     if testScore >= testScoreGrade5 {
      fmt.Println("Top mark")
     } else if testScore >= testScoreGrade4 {
      fmt.Println("Pass with distinction")
     } else if testScore >= testScoreGrade3 {
      fmt.Println("Pass with distinction")
     } else {
      fmt.Println("Failed")
     }
    
     if hasGas && hasKeyInIgnition {
      fmt.Println("Can drive car")
     }
    
     if hasBurger || hasSandwich {
      fmt.Println("Can eat")
     }
    
     if !hasSandwich {
      fmt.Println("No sandwiches, then I will starve, I only eat sandwiches")
     }
    }
   ```

1. Run the command `go run main.go`, to run the program

   ```bash
   goo run main.go
   ```

   You should see the following output:

   ```output
    Message
    Failed
    Can drive car
    Can eat
    No sandwiches, then I will starve, I only eat sandwiches
   ```

1. Try playing around with the code, how does the output change if you change `testScore` value to 51, 62 03 90?
1. Challenge: a test score probably shouldn't be negative, how can you add a check for that?
  
## Summary

In this article, you were taught about `if`, `else if` and `else` and the use of boolean expressions. Depending on the value of your data, the code executes differently.
