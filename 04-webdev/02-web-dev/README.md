Web APIs is ususally what we interact with to serve data to our services or to a client. 

## Web API

Common responsibilities for a web services are to respond to requests:

- **asking for data** and serve data like JSON, XML images, CSS, HTML
- **asking to modify a resource** either by creating, updating or deleting it.

## The `net/http` library

There's a library `net/http` that will help us build a web server. Building a web server with this library involves the following:

- Create a server instance
- Define route requests and how to respond to them
- Start the server instance, making sure it's accessible on a certain address and port.

### Create a server instance

In `net/http`,  `http` represents your service instance. 

```go
import (
 "fmt"
 "net/http"
)

func main() {
  // do something with `http`
}
```

### Define routes

A route is you defining logical separations in your app like `products`, `orders` or some other area it makes sense to divide your app in.

To define a route, you define a route pattern and function that is invoked when the route is hit:

```go
func hello(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "hello\n")
}

func main(){
  http.HandleFunc("/hello", hello)
}
```

In the code above, the string "/hello" is a route pattern that states that all web requests to "/hello" should be handled by the `hello()` function.

### Response and Request

A close inspection of the `hello()` function reveals that it takes a `ResponseWriter` and `Request`:

```go
func hello(w http.ResponseWriter, req *http.Request) {
  fmt.c(w, "hello\n")
}
```

The expectation is that you inspect the `req` object, your request for any indata that decides what to return. Then you are to use `w` to produce a response. In this case you are returning a string by passing `w` to `Fprintf()`. `FPrintf()` takes a writer. The writer is anything IO, so it could be, be writing to file for example as well, or as in this case an HTTP response stream.

### Start the server

Ok, we've gone through routes, producing a response, how would we get this server activated so it starts responding to requests?

You use the `ListenAndServe()` function that takes a port like so:

```go
http.ListenAndServe(":8090", nil)
```

## Responding to a request

An incoming request could be asking for a specific route like /products or /orders for example or it could be asking for a specific static file like an image, a text file or maybe CSS. The request itself gives us a hint, about not only the logical domain it wants data from, like orders or products, but what data type it wants or it might even present credentials for authentication. The hint is known as headers.

### Header

There's a concept called headers. A header is giving off a piece of information that could say what piece of content it is, how big the content is or it could be a token helping you authenticate for example.

Headers can exist both on the incoming request as well as the response.

### Serving various types of content

Serving different types of content means that we are working on the response. To serve various content type, we need to instruct the response what type of content it is so that a consuming client knows how to interpret it, (in some cases, clients like a web browser is able to figure that out anyway through a process called content sniffing).

To serve a specific type of content, there's two things you need to do:

- **set the content type**, you set the content type by calling:

   ```go
   w.Header().Set("Content-Type", "image/jpeg")
   ```

   Here the content type is an image of subtype jpeg. There's many content types you could be setting like plain text, CSS, JSON, XML and more. 

- **produce the response**. Producing a response means writing to the response stream. That can be done by calling the `Write()` method on the `ResponseWriter` instance we are passed when we handle a route. There's other methods capable of writing to said stream as well.

### Serving image data

To serve an image, you need to load it into memory, set the content type and write it to the response stream like below code:

```go
func GetImage(w http.ResponseWriter, r *http.Request) {
    f, _ := os.Open("/image.jpg")
    
    // Read the entire JPG file into memory.
    reader := bufio.NewReader(f)
    content, _ := ioutil.ReadAll(reader)
    
    // Set the Content Type header.
    w.Header().Set("Content-Type", "image/jpeg")
    
    // Write image to the response.
    w.Write(content)
}
```

- First we open the image:

   ```go
   f, _ := os.Open("/image.jpg")
   ```

- Secondly, we read the file into memory:

   ```go
   reader := bufio.NewReader(f)
   content, _ := ioutil.ReadAll(reader)
   ```

- Thirdly, set the `Content-Type` header and tell it it's a JPEG image, with the value "image/jpeg":

   ```go
   w.Header().Set("Content-Type", "image/jpeg")
   ```

- Finally, we write the content to the response:

   ```go
   w.Write(content)
   ```

### Serving JSON data

Just like with serving images, we need to follow a similar approach of configuring the correct content type header and then construct the response. Here's the code:

```go
package main

import (
  "encoding/json"
)

type Person struct {
  Id int
  Name string
}

func ReturnJson(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  p := Person {
    Id: 1
    Name: "a person"
  }

  json.NewEncoder(w).Encode(p)
}

func main() {
  http.HandleFunc("/json", ReturnJson)
}
```

- First, we set the content type, by setting the value "application/json":
 
   ```go
   w.Header().Set("Content-Type", "application/json")
   ```

- Secondly, we construct the data we are about to send out:

   ```go
   p := Person {
    Id: 1
    Name: "a person"
  }
   ```

- Finally, we encode the data as JSON and write it to the response stream:

   ```go
   json.NewEncoder(w).Encode(p)
   ```

It's also possible to use the `Marshal()` function like so, instead of `json.NewEncoder()`:

```go
data, err := json.Marshal(p)
w.Write(data)
```

## Working with the request

There's various ways except for headers to instruct the server program what to do:

- **HTTP verb**, the HTTP verb expresses intention. The POST verb means to create a resource and the GET verb says to only read the data for example. There are many HTTP verbs that we will cover later in this chapter. These two below requests means different things:

   ```text
   GET /products # fetching a list of products
   POST /products # creating a new product resource
   ```

- **body**, The body can contain a payload, data we can use to create or update a resource usually. Here's an example:

   ```json
   {
     "name" : "a new product" 
   }
   ``` 

- **router parameters**. As part of a route request, you can have parameters that carry meaning. If the client asks for the route `/products/5` then the 5 can mean the calling client is after a specific product which unique identifier is 5.
- **query parameter**. At the end of the route, there can be a query section. That section can give further instruction to the request to for example reduce the size of the returning data. The query part starts with a question mark, ? and is followed by key value pairs separated by ampersands, &. It can look like so: `/products?page=1&pageSize=20`  

### Parsing a body

The request has a `Body` property. Depending on what's in the body, you might need to decode it. Below code is decoding a piece of JSON and writes it to the response stream:

```go
package main

import (
  "fmt"
  "encoding/json"
)

type Person struct {
  Id int
  Name string
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
  var p Person

  json.NewDecoder(r.Body).Decode(&p)
  // save person to storage

  fmt.Fprintf(w)
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/person/", handleRequest)

  err := http.ListenAndServe(":4000", mux)
}
```

### Read a route parameter

There's no built-in way to access a route parameter so you would have to parse it like so:

```go
tokens := strings.Split(r.URL.Path, "/")
// check each part
```

or use for example a regular expression to parse out the parts.

Another choice is using a library like the following:

- [httprouter](https://github.com/julienschmidt/httprouter)
- [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)

Here's an example using `httprouter`:

```go
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "The id us, %s!\n", ps.ByName("id"))
}

func main() {
    router := httprouter.New()
    router.GET("/products/:id", Hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}
```

### Read a query parameter

The query part of the route is accessible via the `Query()` function on the `URL` property of the request instance:

```go
r.URL // /products?page=3&pageSize=20
r.URL.Query() // ?page=3&pageSize=20
```

To access a specific parameter know that the `Query()` function returns a `Values` map.

```go
r.URL.Query()["page"] // 3
```

It's possible to call the `Get()` function as well, but only if there is only one parameter:

```go
r.URL.Query().Get("page")
```

### HTTP method

The method means different things and should be handled differently. To access the request method, there's a `Method` property on the request, `r`.

```go
r.Method
```

There's also defined constant like `MethodGet`, `MethodPost` on `http`, so you could write code like so:

```go
func handleRoute(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodGet {
    fmt.Println("It's a GET request") 
  }
}
```

## ServeMux, a better way

So far, you've created an HTTP server by calling `ListenAndServe()` with a port argument and nil. But there's another way to do it. You could be using something called servemux. A servemux is also known as a router. Much like using the `http` directly to add routes, you instead add those routes on the servemux. Let's show some code:

```go
mux := http.NewServeMux()
mux.Handle("/hello", handleHello)
http.ListenAndServe(":8090", mux)
```

In the preceding code, you instantiate the servemux by calling `NewServeMux()`. Then, you set up a route and its handler by calling `Handle()`. Finally, you call `ListenAndServe()` but this time around you pass the `mux` instance instead of `nil`.

So how is this better than the other way we've used so far? The first way we learned about, uses a `DefaultServeMux` and risks exposing profiling endpoints, which is bad. Another reason is connecting the routes directly to `http` changes the global state, which is looked down upon in Go generally.

## Assignment - build a first web app

Your web app should have at least one route. Said route should write to the response stream. The web app should start at a specific port, for example 5500.

## Solution

```go
package main

import (
  "fmt"
  "net/http"
)

func handleRequest(w http.ResponseWrite, r *http.Request) {
  fmt.Fprintf(w, "Hi there")
}

func main() {
  http.HandleFunc("/hello", handleRequest) 
  http.ListenAndServe(":8090", nil)
}
```

## Challenge

- list details on the request such as the route, the verb used and the query parameters.
- See if you can serve up different types of data like JSON or images.


