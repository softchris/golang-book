# Your first program

This lesson covers some history of Go and also teaches you how to build your first Go app.

> Watch the video
> [![your first Go program](https://img.youtube.com/vi/1825FjiewWs/0.jpg)](https://www.youtube.com/watch?v=1825FjiewWs)

## Introduction

In this lesson we'll cover:

- The history of Go
- Why use Go for your apps
- The anatomy of a Go app
- Authoring and running your first app

## A history of Go

The language is called Go but is sometimes known as Golang as the first website for it was golang.org.

Go was created in 2009 by Robert Griesemer, Rob Pike and Ken Thompson. It's hard to estimate the number of Go developers but it's somewhere between 1.1 and 2.7 million, quite a sizeable amount. More than 2500 companies are using Go including, Google, Pinterest and Uber. So, you see, used by a lot of folks by big companies.

> Why was Go created?

As is often the case, a programming language is created to deal with the shortcomings of other languages. In this case, the creators wanted this new language to have the following capabilities:

- **Static typing** and run-time efficiency from C.
- **Readability** from JavaScript and Python.
- **High-performance** networking and multi-processing.

It seems the creators agreed on disliking C++ :)

## What is it used for though?

Here's some areas where you are likely to find a Go being used:

- Cloud based and server-side apps.
- DevOps, automation.
- Command-line tools.
- AI and data science.

## References

There are many great resources out there for learning the Go programming language like:

- <https://go.dev/>
- <https://www.tutorialspoint.com/go/index.htm>
- <https://gobyexample.com/>
- <https://www.w3schools.com/go/>
- <https://docs.microsoft.com/en-us/learn/modules/go-get-started/>
- <https://docs.microsoft.com/en-us/learn/modules/serverless-go/>

## Features

So, what features makes Go compelling? Well, there are some features worth mentioning:

- **Static typing**, I like my types :)
- **Package system**. You can consume and create your own packages. Go to [pkg.go.dev](https://pkg.go.dev/) to read more on what packages there are.
- **Command-line tools**, there's a set of executables that are installed when you install Go. With these executables, you can run, build, install packages, run tests and much more.
- **Standard library**. Go has a powerful standard library that will help you with most things you might need. You can read more about what's in the [standard library](https://pkg.go.dev/std) here.
- **Built-in testing**. Having a testing library that just works out of the box is something you shouldn't take for granted.
- Concurrency. Go is great at handling concurrency. It uses concepts like goroutines and channels.
- **Garbage collection**. You can read more about that [here](https://medium.com/safetycultureengineering/an-overview-of-memory-management-in-go-9a72ec7c76a8#:~:text=Go%20has%20all%20goroutines%20reach,the%20collector%20to%20run%20simultaneously). I like when I don't have to deal with that myself and just focusing on solving problems.

## Install Go

Ok then, hope you are intrigued at this point and just want to see some code? Of course, you are :)

Make sure you've followed the instructions for installing Go on your machine.

> <https://go.dev/doc/install>

## A Go program

Here's what a first program can look like:

```go
package main

import "fmt"

func main() {
 fmt.Println("hello")
}
```

### The program in detail

- `package main`, the entry point module needs to have this instruction.
- `import "fmt"`, fmt is standard package for input and output.
- `func main`, entry point function, where your program starts.

## Commands

Now that you have a program, there's two things you might want to do:

- **Run it**, to see if it compiles and runs.
- **Create executable**, an executable is no longer Go code but like any executable program on your machine.

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

Congrats, you've created your first Go application.

## Summary

In this article, you learned about the programming language Go, some features it has and how to write your first program.

## 🚀 Challenge

Compare Go to other programming languages, can you list some differences between them?

## Review & Self Study

Select one of the resources below and try do a tutorial.

- <https://go.dev/>
- <https://www.tutorialspoint.com/go/index.htm>
- <https://gobyexample.com/>
- <https://www.w3schools.com/go/>
- <https://docs.microsoft.com/en-us/learn/modules/go-get-started/>
- <https://docs.microsoft.com/en-us/learn/modules/serverless-go/>

## Assignment

Create a file *main.go*. Use the `fmt` library to print out to the console. Remember that your run programs with `go run <my program>.go`.

## Solution

Create a file *main.go*

```go
package main

import "fmt"

func main() {
  fmt.Println("printing to the console")
}
```
