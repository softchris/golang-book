# Better logging with a logging library

Once you start writing code, you realize quite early that you need to print things to the screen as well as sometimes to a file or even a log service. What you want to say is usually what type of logging you want to do.

## Introduction

In this chapter, you will learn the following:

- Why and what to log.
- Using the `log` library.

## Reasons to log

There are many reasons to log, here are some reasons:

- **Information**, there might be a case where you want to provide some type of information that could be of use to the one using the program.
- **Success**. A success message is a little more than just information, it indicates that you succeeded with something.
- **Warning**. When you have a warning, something happened that you should be aware of. It's usually not serious enough to shut down the app but it should make you vigilant, it could be that memory is running low for example.
- **Error**. When you get an error, you tend to end up in a state where it's no longer a wise choice to continue.
- **Performance**. It's common to measure how long something takes, for the sake of improving things this information can be useful.
- **Other**. There are also other reasons why you would log something, usually, that's connected to your business.

## What to log

The general rule is the more you can log the better. Especially if it's an error you want to fix you might want to log things like:

- When it happened
- What happened
- Specific error info

For every case, you want to log, have a log at how the log message will be used, will a team be logging through these logs, and what would help them. See if you can interview someone on that team.

## Using `log`

In general, you want to log in places where things might go wrong such as when you make web requests, work with I/O and so on.

In general, use these as guidelines for when to log:

- **Faulty input**. If the program risks producing a faulty response, there was a problem converting/casting a number or it received an unexpected input for example.
- **Error state**. If the program ends up in a state from which it can't recover, for example, unable to fetch a batch of data from a data source.

You don't want logs on every single line of code.  

### Standard log `Println()`

To produce a standard log message, you can use the `Println()` function in the `log` package. It takes a string and will produce a log message that combines a date, time, and your error message.

Here's some code using `Println()`:

```go
package main

import (
 "fmt"
 "log"
)

func main() {
  log.Println("log message")
}
```

It will produce an output like so:

```output
2022/03/24 12:42:13 log message
```

### Use `Fatal()` for errors

`Println()` produces a normal looking log message with a date, time and message. `Fatal()` is used when you want to end the program. What `Fatal()` does is to print out the message you give it and call `os.Exit(1)`.

Here's how it can be used:

```go
log.Fatal("quit program due to <specify reason>")
```

### Log to a file

If you develop an app, you are likely to run it and keep an eye on the console for what the app prints out.

However, as your app becomes ready for production, you want to make sure that all logs that can be useful to analyse is kept somewhere, either sent to a log service or stored in a file.

That way, you ensure that you can analyze these logs later to understand where things went wrong or if you want to analyze the performance of your program.

To log into a file, you can use the `SetOutput()` function. It takes a file handler as input. Thereby, you can use these three lines of code to log:

```go
f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

defer f.Close()

log.SetOutput(f)
```

- In the first line, you open up a file "testlogfile" and ensure you can append it to it.

   ```go
   f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
   ```

- In the second line, you ensure the file is closed the last thing that happens in the program.

   ```go
   defer f.close()
   ```

- Finally you call `SetOutput()` which ensures all log message are sent to file "testlogfile", and not shown in the console.

   ```go
   log.SetOutput(f)
   ```

## Assignment

In this assignment, you will add the `log` library to your code.

1. Create a file *records.csv* with the following content:

   ```text
   item,quantity

   112, 2
   94, 3
   ```

1. Create a file *app.go* and give it the following content:

   ```go
    package main

    import (
     "fmt"
     "io/ioutil"
     "os"
    )
    
    func ProcessFile(path string) {
     filebuffer, err := ioutil.ReadFile(path)
     if err != nil {
      fmt.Println("Error: ", err)
      os.Exit(1)
     }
     inputdata := string(filebuffer)
     fmt.Println("Do something with input: \n", inputdata)
    }
    
    func main() {
     fileName := "records.csv"
    
     fmt.Printf("processing file '%s' \n", fileName)
     ProcessFile(fileName)
    }
   ```

1. Run the file with `go run`:

   ```go
   go run app.go
   ```

   You should see the following output:

   ```output
   processing file 'records.csv'

   Do something with input:
   item,quantity
   112, 2
   94, 3
   ```

1. Add the `log` package to the import and replace all calls to `fmt` with `log`, like so:

   ```go
    package main

    import (
     "io/ioutil"
     "log"
    )
    
    func ProcessFile(path string) {
     filebuffer, err := ioutil.ReadFile(path)
     if err != nil {
      log.Fatal("Error: ", err)
     }
     inputdata := string(filebuffer)
     log.Print("Do something with input: \n", inputdata)
    }
    
    func main() {
     fileName := "records.csv"
    
     log.Printf("processing file '%s' \n", fileName)
     ProcessFile(fileName)
    }
   ```

   Let's see how the output differs.

1. Run the program with `go run`:

   ```go
   go run main.go
   ```

   your output should be similar to:

   ```output
    2022/03/28 13:57:57 processing file 'records.csv'

    2022/03/28 13:57:57 Do something with input:
    item,quantity
    112, 2
    94, 3
   ```

1. Next, lets change the name of `fileName` to "record.csv", to trigger an error (there's no such file).

   ```go
   fileName := "record.csv"
   ```

1. Now, run the app `go run`:

   ```bash
   go run app.go
   ```

   You should see a similar output:

   ```output
   2022/03/28 14:04:52 processing file 'record.csv'

   2022/03/28 14:04:52 Error: open record.csv: no such file or directory
   exit status 1
   ```

   This time around though you see the program exciting with exit status 1.

   The conclusion is that it's better to rely on the `log` library because you get dates and times, and you type less. But there's more, we can log to a file, let's see how we do that next.

### Log to a file

Someone examining the output of the program is likely to inspect a log file overlooking the terminal. Let's instruct `log` to log to a file instead.

1. At the start of the `main()` function, add the following code:

   ```go
     logFile := "logfile"

     f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    
     if err != nil {
      log.Fatal("Could not log to file: ", logFile)
     }
     defer f.Close()
   ```

   Now you have instructed `log` to write all entries to the file *logfile*.

1. Run the program again, `go run`:

   ```bash
   go run main.go
   ```

   You should now see the following output:

   ```output
   exit status 1
   ```

   All your log entries have moved to *logfile*, let's see what it looks like:

   ```text
   2022/03/28 14:11:24 processing file 'record.csv'

   2022/03/28 14:11:24 Error: open record.csv: no such file or directory
   ```

   Great, now we have all entries in a central place, which should make it easier for us to analyze.

## Solution

```go
package main

import (
 "io/ioutil"
 "log"
 "os"
)

func ProcessFile(path string) {
 filebuffer, err := ioutil.ReadFile(path)
 if err != nil {
  log.Fatal("Error: ", err)
 }
 inputdata := string(filebuffer)
 log.Print("Do something with input: \n", inputdata)
}

func main() {
 fileName := "record.csv"
 logFile := "logfile"

 f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

 if err != nil {
  log.Fatal("Could not log to file: ", logFile)
 }
 defer f.Close()

 log.SetOutput(f)

 log.Printf("processing file '%s' \n", fileName)
 ProcessFile(fileName)
}
```
