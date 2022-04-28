# Working with strings

There are many reasons why we need to work with strings in different ways. Below are some situations that you are likely to encounter and that the `strings` library has a solution for:

- **user input**, or other types of storage may contain special characters.
- **inspection**, does it contain what we need?
- **parsing**, splitting the file until we get what we need.
- **presentation**, sometimes we need to present the text in certain way in a UI for example, like lowercase, uppercase, title case.

## Handling special characters

Lets say we read user input and we want to interpret what we get as a number. To make our program robust, we're ok with the user typing spaces or newline characters. The following input should be allowed:

```text
114
   114
114\n
```

There's three methods of interest to handle such a case for us, namely `Trim()`, `TrimLeft()` and `TrimRight()`. 

- `Trim()`, what it does is remove white space characters from both left and right side.
- `TrimLeft()`.  It removes whitespace and any other special characters from the left side.
- `TrimRight()`.  It removes whitespace and any other special characters from the right side.

All there functions above has as their second parameter a cutset where you specify what character you want to get rid of. You can for example specify to remove space, newline and tab characters like so:

Here's some code that shows all three methods in use:

```go
fmt.Printf("%s , string length %d \n", s, len(s))
 res := strings.Trim(s, " ")
 fmt.Printf("%s , string length %d \n", res, len(res))

 s2 := "   114  "
 fmt.Printf("%s , string length %d \n", s2, len(s2))
 res = strings.TrimLeft(s2, " ")
 fmt.Printf("%s , string length %d \n", res, len(res))

 s3 := "   114  "
 fmt.Printf("%s , string length %d \n", s3, len(s3))
 res = strings.TrimRight(s3, " ")
 fmt.Printf("%s , string length %d \n", res, len(res))
```
The above string has 3 white spaces to the left and two to the right, giving it a total length of 8.

The output of the above code is:

```output
  114   , string length 8 
114 , string length 3 // Trim
   114   , string length 8 
114   , string length 5 // TrimLeft
   114   , string length 8 
   114 , string length 6 //TrimRight
```

## Inspect with `Contains()`

TODO

## Parsing with `Split()`

TODO

## Presentation with `Uppercase()` and `Lowercase()`
## Assignment

TODO

## Solution

TODO
