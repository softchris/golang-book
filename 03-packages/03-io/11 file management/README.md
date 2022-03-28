# Working with files and directories

You will use a combination of the libraries `os`, `io`, `io/ioutil` to work with files and directories.

Working with these libraries implies you are interested in doing the following:

- Read and write to files.
- Create, delete, and copy files and directories.
- Compress files and directories.
- Work with various file formats like JSON, CSV and XML.

## Read and write to files

There are different types of files, text files, files containing images, videos etc. For that reason you might want to read the content differently, either byte by byte or maybe even the entire file in one go, as text.

### Read a text file

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

### Write text to a file

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

### Append to a file

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

### File information

You might want to look at a specific file and find out various details about it. Things that can be interesting are:

- **Name**, maybe youu started from a path looking at a list of file. Getting the filename can be valuable.
- **Size**. Getting the size of disc can be good to know.
- **Permission**. To know what permissions you have tells you want you are able to do with the file, like running it, writing to it and so on.
- **Last modified**. You might have a query looking for newly updated files only. Inspecting the modified date is what you want.
- **Is a directory**. Files and directories are ultimately just files. There's a flag distinguishing a file from a directory.

To get file information, use the `Stat()` function like so:

```go
fileStat, err := os.Stat(path)
fmt.Println("File Name:", fileStat.Name())        // Base name of the file
fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
fmt.Println("Is Directory: ", fileStat.IsDir())   // Abbreviation for Mode().IsDir()
```

## Perform operations on files and directories

TODO, copy, rename, remove, check for existence

## Compress files and directories

TODO

## Setup `godoc`

`godoc` generates documentation for your package.

1. Setup `GOPATH`:

   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

1. Install `godoc`

    ```bash
    go get golang.org/x/tools/cmd/godoc
    ```

1. Register `godoc` path:

   ```bash
   export PATH="$GOPATH/bin:$PATH"
   ```

1. Run `godoc`:

   ```bash
   godoc -http=:6060
   ```

   on `http://localhost:6060` there are now docs. Find your package and you will see it has docs. 

1. Add comment as normal commented out comment, near the code

## links 
<https://www.golangprograms.com/files-directories-examples.html>
