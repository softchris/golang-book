# Go from the beginning - working with loops

> TLDR; this article covers working with loops in Go

You are likely to want repeat a set of instructions. For example, you might have a list of orders where you need to process each order. Or you have a file where you need to read it line by line or there might be some other calculation. Regardless of your situation, you are likely to need a looping construct, so what are your options in Go?

Basically, you are using the `for` loop. There are three major ways yu can use it:

- **increment steps**. With the help of a variable, you define a start point, a condition when to stop and an incrementation step. This is your "classic" for-loop. Here's what it looks like:

   ```go
   for i := 0; i < 10; i++ {
     // run these statements
   } 
   ```

- while. In many programming languages you a have a `while` keyword. Go doesn't have that, but what you can do is to use the for-loop in a similar way. You omit the initialization step and increment step and get this code:

  ```go
  for <condition> {
    // run these statements
  }
  ```

- for each. lastly, you have the `for-each` like construct that operates on an array like sequence. It uses the `range` keyword to function:

   ```go
   for i,s := range array {
     // run these statements
   }
   ```

## The `for` loop

The conventional for-loop has three different parts:

- initialization, here you want to create a variable and assign it a starter value like so:

   ```go
   for i:= 0;
   ```

   Note the use of `;`, you usually don's use semicolon, but for this construct, you need it.

- condition. The next step is evaluating whether you should continue incrementing or not. You define a boolean expression here, that if `true`, continues to loop:

   ```go
   for i := 0; i< 10     
   ```

   `i < 10`, as long as a value is between 0 and 10 (becomes 10, then loop breaks), then it returns true and the loop continues.

- increment, in this step, the loop variable `i` is updated, updating it by 1 is common but you can add any increment size, negative or positive.

   ```go
   for i := 0; i< 10; i++ {

   }
   ``` 

   Here, `i` is updated by 1. This loop will run ten times.

## Repeat until condition is met with `while`

A simplified version of this loop can omit the initialization and increment step. You are then left with the condition step only. This step tests whether a variable is true or false and the loop exits on false. Here's an example:

```go
i := 1
for i < 10 {
  i++
  // do something
}
```

In this case, we are declaring `i` outside of the loop. Within the loop, we need to change the value to something that will make the loop expression evaluate to false or it will loop forever.

Here's another code, using the same idea, but this time we ask for input from the user:

```go
var keepGoing = true
 answer := ""
 for keepGoing {
  fmt.Println("Type command: ")
  fmt.Scan(&answer)
  if answer == "quit" {
   keepGoing = false
  }
 }
 fmt.Println("program exit")
```

An example run of the program could look like so:

```bash
Type command: test
Type command: something
Type command: quit
program exit
```

## Using `for each` over a range

For this next loop construct, the idea is to operate over an array or some kind of known sequence. For each iteration you are able to get the index as as well as the next item in the loop. Here's some example code:

```go
arr := []string{"arg1", "arg2", "arg3"}
 for i, s := range arr {
  fmt.Printf("index: %d, item: %s \n", i, s)
 }
```

`arr` is defined as an array and then the `range` construct is used to loop over the array. For each iteration, the current index is being assigned to `i` and the array item is assigned to `s`. An output of the above code will look like so:

```output
index: 0, item: arg1
index: 1, item: arg2
index: 2, item: arg3
```

## Controlling the loop with `continue` and `break`

So far, you've seen three ways you can use the `for` construct. There are also ways to control the loop. You can control the loop with the following keywords:

- `break`, this will exit the loop construct

   ```go
   arr = []int{-1,2}
   for i := 0; i< 2; i++ {
     fmt.Println(arr[i])
     if arr[i] < 0 {
       break;
     }
   }     
   ```

   The output will be:

   ```output
   -1
   ```

   it won't output `2` as the loop exits after the first iteration.

- `continue`, this will move on to the next iteration. If `break` exits the loop, `continue` skips the current iteration:

   ```go
   arr = []int{-1,2,-1, 3}
   for i := 0; i< 4; i++ {
     if arr[i] < 0 {
       break;
     }
     fmt.Println(arr[i])
   } 
   ```

   The output will be:

   ```output
   2
   3
   ```

## Summary