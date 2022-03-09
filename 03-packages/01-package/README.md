# Go modules

There are two interesting use cases with modules:

- **Consuming a module**, you will most likely use a combination oof core moodules and external 3rd party modules
- **Creating a module**, in some case you will create code that you or someone else will be able to use. For this scenario, you can create a module and upload it to GitHub.

## Consume internal files

You might want to split up your app in many different files. Let's say you have the following files:

```output
/app
  main.go
  /helper
    helper.go 
```

1. Create the module like so:

    ```go
    go mod init
    ```

1. create the **helper** directory and *helper.go* file and give it the following content:

    ```go
    // helper.go
    
    package helper
    
    import "fmt"
    
    func Help() {
     fmt.Println("This is a helper function")
    }
    ```

1. Create the *main.go* file and give it the following content:

   ```go
   package main

   import (
     "log-tester/helper"
    
     "github.com/softchris/logger"
   )
    
   func main() {
     logger.Log("hey there")
     helper.Help()
   }
   ```

   Note this import `"log-tester/helper"`, it ensures the `helper` package is in scope.

1. Compile and run

   ```bash
   go run main.go
   ```

## Consume an external module

To consume an external module, you need to:

- **Import it**, this involves using the `import` instruction and fully qualifying the address to the module's location.
- **Use it in code**, call the code from the module that you mean to use
- **Ensure it's downloaded**, so your code can be run.

### Create a project

Let's create a new project

1. Run `go mod init`:

    ```go
    go mod init hello
    ```

    Note how *go.mod* was created with the following content:

    ```go
    module hello
    
    go 1.16
    ```

### Add reference to an external lib

Next, lets create some code that will use the external library:

1. Create the file *main.go*

    ```go
    package main
    
    import (
      "fmt"
      "github.com/softchris/math"
    )
    ```

1. To the same file, add `main()` function and use external function, from package:

    ```go
    func main() {
      sum += math.Add(1,2)
      fmt.Println(sum)
    }
    ```

## Fetch the lib

Now, we need to resolve the external library. 

1. Run `go mod tidy`:
    
    ```bash
    go mod tidy
    ```

    Your *go.mod* is updated:

    ```go
    require github.com/softchris/math v0.2.0
    ```

    There's also *go.sum* file with the following content:

    ```go
    github.com/softchris/math v0.2.0 h1:88L6PLRBGygS3LY5KGIJhyw9IZturmd2ksU+p13OPa4=

    github.com/softchris/math v0.2.0/go.mod h1:v8WzhjKC+ipuH+i9IZ0Ta2IArniTP53gc5TgCINCqAo=
    ```

    This is Go's way of keeping track of how to build the app by referencing to the go module in question.

1. Run `go run`:

   ```go
   go run main.go
   ```

   ```output
   3
   ```

## Create a module

When you build a module meant for sharing, there's some gotchas:

- you need to create a package
- your package won't be called main
- there's concept of public and private parts of your code
- you can't test it locally
- upload your package to GitHub for wide distribution

### Create the code for the module

To create a moodule meant for wider use you need to first initialize a module.

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

   At this point, you are consuming the "logger" package but it's pointing to GitHub and your package don't live there yet. However, you can repoint to a local address on your machine, lets do that next.

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

1. Create the repo at GitHub.

1. Make you do at least one commit:

    ```bash
    git add .
    git commit -m "first commit"
    ```

1. Do the following do upload your package to GitHub:

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
   require github.com/softchris/logger v0.1.0
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

Great, you've managed to create a module locally, put it on GitHub and have one of your other Go projects use it.