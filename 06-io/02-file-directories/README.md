# Perform operations on files and directories

Things you likely want to do to files and directories are moving them about, removing them, rename them etc. In short, high-level operations that is less about the content of the file but doing something with it.

## Introduction

In this chapter you will:

- Analyze a file for its metadata like size, modified date and more.
- Copy files from one location to the next.
- Rename files.
- Remove files.
- Create and read directories.

## File information

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

Also when you call `ReadDir()` you get back an array of `FileInfo` objects:

```go
files, err := ioutil.ReadDir(path)
```

## Copy file

Copying a file is really three operations:

- **open** the file at the destination.
- **create** a new file at a given destination.
- **copy**, then transfer the content to that location.

Here's how you can implement a *copy* operation:

```go
// copies 'test.txt' and its content to 'copy.txt'
src := "test.txt"
dest := "copy.txt"

srcFile, err := os.Open(src)
if err != nil {
  log.Fatal(err)
}
defer srcFile.Close()

newFile, err := os.Create(dest)
if err != nil {
  log.Fatal(err)
}
defer newFile.Close()

_, err = io.Copy(newFile, srcFile)
if err != nil {
  log.Fatal(err)
}
```

## Rename

Rename is a bit easier to achieve. The `os` package has a `Rename()` function. Here's how to use it:

```go
err := os.Rename(src, dest)
```

## Remove file

To remove file, there's a `Remove()` function you can use. Here's how to use it:

```go
err := os.Remove(path)
```

## Create dir

To create a directory, you can use the `MkdirAll()` function in the `os` library. However, you should check whether the directory exist first. The way to do that is to use the `IsNotExist()` function like so:

```go
_, err := os.Stat(dirName)

 if os.IsNotExist(err) {
  errDir := os.MkdirAll(dirName, 0755)
  if errDir != nil {
   log.Fatal(err)
  }
 } else if err != nil {
  log.Fatal("error creating dir")
 } else {
  log.Fatal("directory exist")
 }
```

As you can see on the above code:

1. you first use the `Stat()` function. The `Stat()` returns a `FileInfo` object or an error of type `PathError` if path doesn't exist.

1. `os.IsNotExist(err)`. This returns `true` if `err` is a `PathError`, i.e the path don't exist, which is good, we want to create it.

1. Finally, we call `os.MkdirAll(dirName, 0755)`. The 755 instruction is about permissions on the created directory, which gives the permissions, Read/Write/Execute, Read/Execute, Read/Execute. 755 is a common permission set on web servers. You essentially want to avoid anyone but you to modify the file.

## Read dir

Reading a directory is quite straight forward. You can use the `ReadDir()` function on the `io/ioutil` library. Here's how you would read a directory:

```go
files, err := ioutil.ReadDir(path)
```

`files` is an array of type `FileInfo`.  

TODO, copy, rename, remove, check for existence

## Assignment

Create the following files and directories like so:

```text
tmp/
  a.txt
  b.xt
  subdir/
```

Your program should read the diretory *tmp* and for each entry list, if it's a dir or a file, the size and when last modified.

The programs output should look something like:

```output
Reading directory tmp:
Name, Type, Size, Modified
a.txt, file, 1kb, 2022-01-01
b.txt, file, 1kb, 2022-01-01
subdir, directory, 1kb, 2022-01-01
```

## Solution

```go
package main

import (
 "fmt"
 "io/ioutil"
 "log"
)

func GetType(isDir bool) string {
 if isDir {
  return "directory"
 }
 return "file"
}

func main() {
 var path string = "tmp"
 files, err := ioutil.ReadDir(path)
 if err != nil {
  log.Fatal(err)
 }
 fmt.Println("Reading directory ", path)
 fmt.Println("Name, Type, Size, Modified")
 for _, file := range files {
  if err != nil {
   log.Fatal(err)
  }
  fmt.Printf("File Name: %s, ", file.Name())
  fmt.Printf("Type: %s, ", GetType(file.IsDir()))
  fmt.Printf("Size: %d, ", file.Size())             
  fmt.Printf("Last Modified: %s, ", file.ModTime()) 
  fmt.Print("\n")
 }
}

```
