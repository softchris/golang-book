# Read and write to files

There are different types of files, text files, files containing images, videos etc. For that reason you might want to read the content differently, either byte by byte or maybe even the entire file in one go, as text.

## Introduction

In this chapter you will learn to:

- Read content from text files.
- Write to text files.
- Append text at the end of a text file.
- Analyze a file for its metadata like size, modified date and more.

## Read a text file

One approach could be using the `ioutil` library and its `ReadFile()` method like so:

```go
  import (
    "io/ioutil"
    "log"
  )
  func main() {
    filebuffer, err := ioutil.ReadFile(path)
 
    if err != nil {
      log.Fatal(err)
    }
    var inputdata string = string(filebuffer) 
  }
```

Note how the result of reading the file ends up in `filebuffer`. To interpret it as a string, you need to convert it via `string(filebuffer)`. Now, you're ready to process the file content, read it line by line or whatever you want to do.

## Write text to a file

In this scenario, we are looking to do two things:

- Create a file.
- Write text to our newly created file.

For this scenario, we can use the `os` library and a combination of the `Create()` method, to create a file and the `WriteString()` method to write a string to the file.

```go
 import (
  "os"
  "log"
 ) 

 f, err := os.Create(path)
 if err != nil {
  log.Fatal(err)
 }

 n, err := f.WriteString(content + "\n")
 if err != nil {
  log.Fatal(err)
 }
 fmt.Printf("wrote %d bytes\n", n)
 f.Sync()
```

First the file is created calling `Create()`. From that, we get a file handle `f`. With `f`, we can call `WriteString()` with a string. Lastly, we call `Sync()` to ensure the string is persisted in the file.

## Append to a file

Appending text, implies you already have an existing file. When you append, you information to the end of the file. Appending is something you are likely to do when you add new entries to a log file or adding a new purchase to a Point of Sale, POS system in a grocery store for example.

To append to a file you use the `OpenFile()` method in the `os` lib. What you need to do is to pass it some flags that states that you want to append content. You should also have a behavior that says, create if it doesn't already exist. You end up with code looking like so:

```go
f, err := os.OpenFile("text.log",
 os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
 log.Println(err)
}
defer f.Close()
```

The 0644, is a 3x3 bit flag. It sets permissing for User (6, Read/Write), Group (4, Read), and Other (4, Read).

To append a string, you can call `WriteString()` like so:

```go
f.WriteString("my text \n")
```

## Assignment

Imagine you have a file *invoices.csv* looking like so:

```text
customer, amount, date
Wood LTD, 100, 2020-01-01
Metal, 345, 2020-01-29
Steel, 700, 2020-07-29
```

Open up and read the file content, line by line.

## Solution

```go
package main

import (
 "fmt"
 "io/ioutil"
 "log"
 "strings"
)

func main() {
 var path = "invoices.csv"
 filebuffer, err := ioutil.ReadFile(path)

 if err != nil {
  log.Fatal(err)
 }
 var inputdata string = string(filebuffer)

 rows := strings.Split(inputdata, "\n")
 for _, row := range rows {
  fmt.Println("row:", row)
 }
}

```

## Challenge

See if you can split up each row further in columns and summarize how much the orders are worth in total.