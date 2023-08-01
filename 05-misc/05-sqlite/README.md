# Working with a database

You will be working with a sqlite database and read and write values to it.

## Introduction

In this chapter you will:

- Create a database and its structure.
- Write to the database.
- Read from the database.

## Select a sqlite driver

To connect with a sqlite database we've got a few libraries to choose from. These libraries will provide you with a sqlite driver that you need to successfully connect with your database. Here are some common ones:

- SQLite (uses cgo): <https://github.com/mattn/go-sqlite3> [*]
- SQLite (uses cgo): <https://github.com/gwenn/gosqlite> - Supports SQLite dynamic data typing
- SQLite (uses cgo): <https://github.com/mxk/go-sqlite>
- SQLite: (uses cgo): <https://github.com/rsc/sqlite>
- SQLite: (pure go): <https://modernc.org/sqlite>

Refer to this link to see libraries for other databases:

- <https://github.com/golang/go/wiki/SQLDrivers>
.

## Use `sqlite3` from the console

To work with your database, it's beneficial to use sqlite from the command line. Consult the official [downloads page](https://www.sqlite.org/download.html) for sqlite and ensure you pick the executable for your operating system.  Installing sqlite will give you an executable.

With the executable, you can:

- Create a database.
- Run SQL commands.
- Run other commands supported by the executable.

### Create a database

To create a database, you give it the name of the database like so:

```console
sqlite3 activities.db
```

The preceding command will give you a database in a file called "activities.db"

> It will also start up a shell where you can execute SQL commands as well as commands supported by sqlite3.

### Run SQL commands

Run a SQL command in the shell by typing SQL and end it by a semicolon, `;`.

```sql
CREATE TABLE `person` (
        `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
        `name` VARCHAR(64) NULL,
        `lastname` VARCHAR(64) NULL,
        `created` DATE NULL
    );
```

### Exit the shell

After you're done, you can exit the shell by typing `.exit`:

```console
.exit
```

## Talking to your database via Go

To talk to your database via Go, there's some steps you need to take in order:

1. **Create a project**. You need to create a project so you can import a Go package containing your sqlite driver. Create a project by running `go mod init`. Below is an example:

   ```console
   go mod init "example-project"
   ```

1. **Add imports**. Once you have the needed packages you need to refer to them in the import section:

   ```go
   import (
     "database/sql"
     _ "github.com/mattn/go-sqlite3"
    )
   ```

   Above, are the two packages we will use, "database/sql" that provides an interface for us to run queries and statements. "github.com/mattn/go-sqlite3" contains the driver that will enable us connecting to the database.

1. **Establish connection**. To connect with the database, you call the `Open()` function on the `sql` instance like so:

   ```go
   db, err := sql.Open("sqlite3", "./mydb.db")
   ```

   In the preceding command, we specify first the type of database and in the second parameter the name of the database and where it's located. We get a database instance back of type `*sql.DB`.

1. **Run queries**. At this point, we are free to run queries. You use `Query()` function and give it a SQL statement like in this example:

   ```go
   rows, err := db.Query("SELECT * FROM person")   
   ```

   To iterate over the results, you can use a for-loop like so:

   ```go
   for rows.Next() {
      var uid int
      var name string
      var lastname string
      var created time.Time
      err = rows.Scan(&uid, &name, &lastname, &created)
   }
   ```

1. **Run prepared statements**. Prepared statements are SQL statements where we can provide parameter values at a later point. You call the `Prepare()` function with `?` as placeholders where there will be data inserted:

   ```go
   stmt, err := db.Prepare("UPDATE person set lastname=? where uid=?")
   ```

   To run the statement against the database, you can call `Exec()`:

   ```go
   res, err := stmt.Exec("smith", 1)
   ```

   The `res` instance coming back has a function `RowsAffected()` that returns the number of affected rows:

   ```go
   affected, _ := res.RowsAffected()
   ```

   Getting affected rows is a good indicator that you actually changed something.

1. Close the database connection. You should close the database when you're done with it like so:

   ```go
   db.Close()
   ```

## Assignment

In this assignment, we will create a Go program that's able to read and write to the database. We will go all the way from creating the database with the console to writing the Go code needed.

### Create the database and populate it

We will create our database using the sqlite executable in the console.

1. Run `sqlite` to create the database and initialize the sqlite shell:

    ```console
    sqlite3 mydb.db
    sqlite3SQLite version 3.32.3 2020-06-18 14:16:19
    Enter ".help" for usage hints.
    ```

    At this point, you have a database created. Next, we need some tables in there.

1. Run the following SQL command in the sqlite shell:

   ```sql
   CREATE TABLE `person` (
            `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
            `name` VARCHAR(64) NULL,
            `lastname` VARCHAR(64) NULL,
            `created` DATE NULL
        );
   ```

   You know have the table "person" created. Next, we need some data in the table that we will interact with later in our Go code.

1. Run this SQL command to insert data into "person" table:

   ```sql
    insert into person(name,lastname, created) values ("joe", "schmoo", '2021-01-01');
    ```

    Great, we now have data in our table. Time to focus on the Go code next.

1. Run `.exit` to exit the database.

### Create a project

Now we will create a Go project and some code able to access our database.

1. Create *db.go* and give it this content:

    ```go
    package main
    
    import (
     "database/sql"
     "fmt"
     "log"
     _ "github.com/mattn/go-sqlite3"
    )
    
    func main() {
     db, err := sql.Open("sqlite3", "./mydb.db")
     if err != nil {
      log.Fatal(err)
     }
     fmt.Println("database open")
    
     fmt.Println("bye")
    
     fmt.Println("closing db")
     db.Close()
    
    }
    ```

    Next, lets initialize our Go project.

1. Run the following commands to create our project:

   ```console
   go mod init sql-demo
   ```

1. Run `go mo tidy`, to install the needed packages you specified in the import section of your program (this will download and add "github.com/mattn/go-sqlite3" to your project):

   ```console
   go mod tidy
   ```

### Read data

Next, we will add a function capable of reading data.

1. Add a function `Read()` like so:

   ```go
   func Read(db *sql.DB) {
     rows, err := db.Query("SELECT * FROM person")
     
    }
   ```

   At this point, we have read the response into `rows`. Next, we need iterate on the response.

1. Add the following code, in `Read()` to iterate over the response:

   ```go
   for rows.Next() {
      var uid int
      var name string
      var lastname string
      var created time.Time
      err = rows.Scan(&uid, &name, &lastname, &created)
      if err != nil {
       log.Fatal(err)
      }
      fmt.Println(uid)
      fmt.Println(name)
      fmt.Println(lastname)
      fmt.Println(created)
   }
   ```

   Note the usage of `Scan()` and variables being sent in as references so the response is written to them.

### Create data

Now we will create code that will allow us to create data in our database.

1. Add a function `Read()`:

   ```go
   func Create(db *sql.DB) {
     stmt, err := db.Prepare("INSERT INTO person(name, lastname, created) values(?,?,?)")
     
   }
   ```

   At this point, you have created a statement that when executed will attempt to insert row. Note the use of `?`, these are placeholders that you will need to provide values to at the moment of execution.

1. Add the following code to `Create()`:

   ```go
   if err != nil {
      log.Fatal(err)
     }
     res, err := stmt.Exec("Mrs", "Smith", "2022-01-01")
     if err != nil {
      log.Fatal(err)
     }
     affected, _ := res.RowsAffected()
     log.Printf("Affected rows %d", affected)
   ```

   Note the call to `Exec()`, here you are providing  data and `?` is being replaced by the values you send in. Also note the last two rows:

   ```go
   affected, _ := res.RowsAffected()
   log.Printf("Affected rows %d", affected)
   ```

   Here for result `res` we are invoking `RowsAffected()` that returns the number of affected rows then we go on to print said value.

### Update and delete data

Updating and deleting data takes on very similar approach to how to create data. You will use a statement that you prepare and then send in the real values. Below is the code for performing both these actions:

**Update**

```go
func Update(db *sql.DB) {
 stmt, err := db.Prepare("UPDATE person set lastname=? where uid=?")
 if err != nil {
  log.Fatal(err)
 }
 res, err := stmt.Exec("smith", 1)
 if err != nil {
  log.Fatal(err)
 }
 affected, _ := res.RowsAffected()
 log.Printf("Affected rows %d", affected)
}
```

**Delete**

```go
func Delete(db *sql.DB) {
 stmt, err := db.Prepare("delete from person where uid=?")
 if err != nil {
  log.Fatal(err)
 }
 res, err := stmt.Exec(1)
 if err != nil {
  log.Fatal(err)
 }
 affected, _ := res.RowsAffected()
 log.Printf("Affected rows %d", affected)
}
```

## Solution

```go
package main

import (
 "database/sql"
 "fmt"
 "log"
 "time"

 _ "github.com/mattn/go-sqlite3"
)

func Read(db *sql.DB) {
 rows, err := db.Query("SELECT * FROM person")
 if err != nil {
  log.Fatal(err)
 }
 for rows.Next() {
  var uid int
  var name string
  var lastname string
  var created time.Time
  err = rows.Scan(&uid, &name, &lastname, &created)
  if err != nil {
   log.Fatal(err)
  }
  fmt.Println(uid)
  fmt.Println(name)
  fmt.Println(lastname)
  fmt.Println(created)
 }
}

func Update(db *sql.DB) {
 stmt, err := db.Prepare("UPDATE person set lastname=? where uid=?")
 if err != nil {
  log.Fatal(err)
 }
 res, err := stmt.Exec("smith", 1)
 if err != nil {
  log.Fatal(err)
 }
 affected, _ := res.RowsAffected()
 log.Printf("Affected rows %d", affected)
}

func Create(db *sql.DB) {
 stmt, err := db.Prepare("INSERT INTO person(name, lastname, created) values(?,?,?)")
 if err != nil {
  log.Fatal(err)
 }
 res, err := stmt.Exec("Mrs", "Smith", "2022-01-01")
 if err != nil {
  log.Fatal(err)
 }
 affected, _ := res.RowsAffected()
 log.Printf("Affected rows %d", affected)
}

func Delete(db *sql.DB) {
 stmt, err := db.Prepare("delete from person where uid=?")
 if err != nil {
  log.Fatal(err)
 }
 res, err := stmt.Exec(1)
 if err != nil {
  log.Fatal(err)
 }
 affected, _ := res.RowsAffected()
 log.Printf("Affected rows %d", affected)
}

func main() {
 db, err := sql.Open("sqlite3", "./mydb.db")
 if err != nil {
  log.Fatal(err)
 }
 fmt.Println("database open")
 Create(db)
 Read(db)
 // Update(db)

 fmt.Println("bye")

 fmt.Println("closing db")
 db.Close()

}
```
