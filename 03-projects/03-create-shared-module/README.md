# Create a module meant for sharing

In this chapter, we will cover how you can create a module you can share with others.

## Introduction

This chapter will cover:

- Creating a project.
- Testing the module locally.
- Tag the module with different versions.
- Try consuming the module as an external library.

## Create a module

When you build a module meant for sharing, there's some gotchas:

- You need to create a package.
- Your package won't be called main.
- There's the concept of public and private parts of your code.
- You can test it locally.
- Upload your package to GitHub for wide distribution.

## Assignment - create a module meant for sharing and consume it

To create a module meant for wider use you need to first initialize a module.

1. Create a directory *logger* for your new package:

   ```go
   mkdir logger
   cd logger   
   ```

1. Run `go mod init <address at github>`, for example:

   ```bash
   go mod init github.com/softchris/logger
   ```

   This will create a *go.mod* file in your directory.

   ```output
   logger/
     go.mod
   ```

   The file looks like so:

   ```go
   module github.com/softchris/logger

   go 1.16
   ```

   It contains the package name and the version of Go it means to use.

1. Create a file to host your package code, for example *log.go* and give it the following content:

   ```go
    package logger

    import (
     "fmt"
    )
    
    var Version string = "1.0"
    
    func Log(mess string) {
     fmt.Println("[LOG] " + mess)
    }
   ```

   - Note `package logger` instead of `main`.
   - The uppercase variables and methods makes the publicly available. Anything named with lowercase will be private for the package.

### Test it locally

You can test your package locally. To do so you need a separate package that you can import your package from.

1. Move up a directory:

   ```bash
   cd ..
   ```

1. Create a new directory **logger-test**:

   ```bash
   mkdir logger-test
   cd logger-test
   ```

1. Create a package, it will be used for testing only:

   ```bash
   go mod init logger-test
   ```

1. Create a file *main.go* and add the following code:

   ```go
    package main

    import "github.com/softchris/logger"
    
    func main() {
     logger.Log("hey there")
    }
   ```

   At this point, you are consuming the "logger" package but it's pointing to GitHub and your package doesn't live there yet. However, you can repoint to a local address on your machine, let's do that next.

1. Open *go.mod* and add the following:

   ```go
   require github.com/softchris/logger v0.0.0

   replace github.com/softchris/logger => ../logger
   ```

   Two things are happening here:
  
   - you are asking for the "logger" package:
  
      ```go
      require github.com/softchris/logger v0.0.0
      ```

   - you are making it point to your local system instead of GitHub

      ```go
      replace github.com/softchris/logger => ../logger
      ```

1. Run the package with `go run`:

    ```bash
    go run main.go
    ```

    You should see:

    ```output
    [LOG] hey there
    ```
  
### Publish a package

To publish your package, you can put it on GitHub.

1. Create a git repo with `git init`:

   ```bash
   git init
   ```

1. Create the repo on GitHub.

1. Make you do at least one commit:

    ```bash
    git add .
    git commit -m "first commit"
    ```

1. Do the following to upload your package to GitHub:

   ```bash
   git remote add origin https://github.com/softchris/logger.git

   git branch -M main
   git push -u origin main
   ```

1. Tag your package with `git tag`:

   ```bash
   git tag v0.1.0
   git push origin v0.1.0
   ```

   Now your package has the tag 0.1.0

### Test it out

1. Go to your project "logger-test":

   ```bash
   cd ..
   cd logger-test
   ```

1. Open up *go.mod* and remove these lines:

   ```go
   require github.com/softchris/logger v0.0.0

   replace github.com/softchris/logger => ../logger
   ```

1. Run `go mod tidy`, this will force Go to go look for the package:

   Your *go.mod* should now contain:

   ```go
   require github.com/softchris/logger v0.1.0
   ```

   Also, your *go.sum* should contain:

   ```go
   github.com/softchris/logger v0.1.0 h1:Kqw7t9C3Y7BtHDLTx/KXEqHy5x8EJxrLian742S0di0=

   github.com/softchris/logger v0.1.0/go.mod h1:rrzWjMsM3tqjetDBDyezI8mFCjGucF/b5RSAqptKF/M=
   ```

1. Run the program with `go run`:

   ```bash
   go run main.go
   ```

   You should see:

   ```output
   [LOG] hey there
   ```

## Challenge

See if you can add a feature to your new package. Give it a new tag via Git. Then ensure your app is using this new version.
