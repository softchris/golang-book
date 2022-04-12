# Go for beginners - working with JSON

In this chapter, you will work with the JSON data format.

##  Pre-Lecture Quiz

TODO

## Introduction

In this chapter, you will learn the following:

- The JSON data format.
- How to read JSON data and map it to existing structures in Go.
- How to create JSON and persist it in files.

## JSON

JSON is lightweight data format for storing and transporting data. The acronym stands for **J**ava**S**cript **O**bject **n**otation.

The format is commonly used in Web services. Usually data is encoded as JSON and sent via HTTP. A client, for example a web browser, consume the JSON data and use it to render a frontend with the help of HTML and CSS. It's also common for JSON to be used as a way to communicate between services.

Here's an example of what the format looks like:

```json
{
  "products": [{
    "id": 1,
    "name": "a product"
  },
  {
    "id": 2,
    "name": "another product"
  }]
}
```

The above depicts a list of products. Each key needs to be encased with quotes and values can be everything from primitives like numbers, strings, booleans etc to more complex types like an array or an object.

what is this format, what contexts is it used in.

## Reading JSON

To work with JSON you need to use the `encoding/json` library. It allows you to both read and write JSON data. 

To read JSON data you need to first have a structure in your Go code read to map to the JSON data. Imagine having the following JSON data:

```json
{
  "products" : [
    {
      "id": 1,
      "name": "some product"
    }
  ]
}
```

To map this in Go, you need a struct that matches its structure like so:

```go
type Product struct {
  Id int
  Name string
}

type Response struct {
  Products []Product
}
```

### Notation

The first step is to create the structures needed to match your JSON data. But, you also need to do one more thing, add notations. The problem we are looking to solve is is how the `Product` property is called `Id` and not `id` as it's called in the JSON data.

> Won't that just work?

Unfortunately not

> Why don't you just name the struct property `id` 

The library `encoding/json`, is separate package from the main package, meaning we need to make a field name, defined in the main package, uppercase for the `encoding/json` package to find it.

So that means, we need to add the following annotations to our above created structs:

```go
type Product struct {
  Id int `json: "id"`
  Name string `json: "name"`
}

type Response struct {
  Products []Product `json: "products"`
}
```

What these annotations do is to say, in the JSON data, look for properties with these names and map it to the following property. Like in this example from above:

```go
Id int `json: "id"`
```

### Reading the data

Ok, so we've defined the structures in Go that we will map our JSON data to. So how do we read from a JSON source? Well, JSON is usually stored in one of two ways:

- a string literal, like so `{ "name": "my product", "id": 1 }`
- in a JSON file:

   ```json
   {
     "id": 1,
     "name": "my product"
   }
   ```

Let's show how to work with both approaches:

**Reading from a string literal**

We will use the `Unmarshal()` function and provide it with the string literal as the first parameter and the data to write the result to as the second parameter.

```go
package main

imports (
  "fmt"
  "encoding/json"
)

func main() {
  str := `{ "name": "my product", "id": 1 }`
  product := Product{}
  json.Unmarshal([]byte(str), &person)
  fmt.Println(person) // prints the object
}
```

Note how we also convert the response to a byte array `[]byte(str)` and how the data is written in the second parameter to the `person` instance as reference, `&person`.

**read from a file**

To read from a file, we will use the `io/ioutil` library and its `ReadFile()` function. Like with the string literal, the `Unmarshal()` function will be used to write the data to a struct instance.

```go
package main

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "iohelper/dir"
 "iohelper/file"
 "log"
)

type Products struct {
 Products []Product `json: products`
}

type Product struct {
 Id   int    `json: "id"`
 Name string `json: "name"`
}

func main(){
  file, _ := ioutil.ReadFile("products.json")

  data := Products{}

  _ = json.Unmarshal([]byte(file), &data)

  for i := 0; i < len(data.Products); i++ {
    fmt.Println("Product Id: ", data.Products[i].Id)
    fmt.Println("Name: ", data.Products[i].Name)
  }
}
```

## Writing JSON

We've seen so far how we can read JSON data either from a string literal or from a file, but what about writing data?

There are two cases that are important for us:

- Writing data to a structure. We are working with structs so any changes we do on the structs needs to be converted back to JSON.
- Generate JSON from data. In the second case, we might be working with a raw string literal or even with pure primitives, how would we do that?

Let's address both these cases. What these cases have in common is the `Marshal()` function. The `Marshal()` method is capable of taking a primitive or a struct and take that into JSON:

```go
package main

import (
 "fmt"
 "encoding/json"
)

type Person struct {
  Id int `json: "id"`
  Name string`json: "name"`
}

func main() {
  aBoolean, _ := json.Marshal(true)
  aString, _ := json.Marshal("a string")
  person := Person{
    Id: 1
    Name: "a person"
  }
  aPerson, _ := json.Marshal(&person)
  fmt.Println(string(aBoolean)) // true
  fmt.Println(string(aString))  // a string
  fmt.Println(string(person))  // { "id": 1, "name": "a person" }
}
```

## Assignment

Given the following file *response.json*, find a way to read the data and display it on the screen:

```json
{
  "orders": [
   {
    "id": 1,
    "items": [
      { "id": 1, "quantity": 3, "total": 34.3 },
     { "id": 2, "quantity": 2, "total": 17.8 }
    ] 
   },
   {
    "id": 2, 
    "items": [
      { "id": 3, "quantity": 3, "total": 10.0 },
      { "id": 4, "quantity": 2, "total": 100.5 }
    ] 
   } 
  ]
}
```

## Solution

```go
package main

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
)

type OrderItem struct {
 Id       int     `json: "id"`
 Quantity int     `json: "quantity"`
 Total    float32 `json: "total"`
}

type Order struct {
 Id    int         `json: "id"`
 Items []OrderItem `json: items`
}

type Response struct {
 Orders []Order `json: orders`
}

func main() {
 file, _ := ioutil.ReadFile("orders.json")

 data := Response{}

 _ = json.Unmarshal([]byte(file), &data)
 fmt.Println(data)
 for i := 0; i < len(data.Orders); i++ {
  fmt.Println("Order Id: ", data.Orders[i].Id)

  for j := 0; j < len(data.Orders[i].Items); j++ {
   item := data.Orders[i].Items[j]
   fmt.Println("Item id", item.Id)
   fmt.Println("Item quantity", item.Quantity)
   fmt.Println("Item total", item.Total)
  }
 }
}

```

## 🚀 Challenge

TODO

## Post-Lecture Quiz

TODO

<https://gobyexample.com/json>
