# Working with files and directories

You will use a combination of the libraries `os`, `io`, `io/ioutil` to work with files and directories.

Working with these libraries implies you are interested in doing the following:

- Read and write to files
- Create, delete, and copy files and directories
- Compress files and directories
- Work with various file formats like JSON, CSV and XML

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

TODO

### Inspect file

TODO

## Perform operations on files and directories

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
