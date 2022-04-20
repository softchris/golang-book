# Read and write to files

There are different types of files, text files, files containing images, videos etc. For that reason you might want to read the content differently, either byte by byte or maybe even the entire file in one go, as text.

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
  }
```

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

## Assignment

TODO

## Solution

TODO
