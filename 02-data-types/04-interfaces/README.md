# Go from the beginning - interfaces

This chapter covers what an interface and what too use it for.

##  Pre-Lecture Quiz

TODO

## Introduction

This chapter will cover:

- What an interface and how it differs from a struct
- How to add behavior
- Implement an interface
- Type assertions
- Changing a value

## Interface

To describe what an interface is, lets start by talking about structs and how they are different from an interface.

With structs, we can define properties we want a concept to have, like for example a car:

```go
type Car struct {
  make string 
  model string
}
```

An interface is meant to communicate something different, a behavior. Instead of describing the car itself, like a struct does, it describes what a car can do.  

## Interface - describing a behavior

Now that we've described how an interface differs from a struct, lets talk about the motivation for using an interface. There are a couple of good reasons for when to use an interface:

- **Adding behavior**. When you want your types to have a behavior, that's when you want an interface
- **Communicate via contract**. Often, when you call other code, you want to reveal as little of your concrete implementation as possible. Instead of saying, here's car, you might want to say, here's something that can run. It enables your code to be flexible and you don't have to implement specific code for each type but can instead write code that deals with a certain behavior. 

### Define an interface

To define an interface, you need the keywords `type` and `interface` and you need a set of methods, one or many that a type should implement. Here's an example interface:

```go
type Desribeable interface {
  describe() string
}
```

Here's another example:

```go
type Point struct {
  x int
  y int
}

type Shape interface {
  area() int
  location() Point
}
```

## Implement an interface

Everything that's a type can implement an interface. More than one type can implement the same interface. Let's look at how a type `Rectangle` can implement the `Shape` interface:

```go
type Rectangle struct {
  x int
  y int
}

func (r Rectangle) area() int {
  return r.x * r.y
}

func (r Rectangle) location() Point {
  return P{ x: r.x, y: r.y }
}
```

So what's going on here? Let's look at the first method `area()`:

```go
func (r Rectangle) area() int {
  return r.x * r.y
}
```

It looks like a regular function but there's this `(r Rectangle)` right before the function name. That's a signal to Go that you are implementing a certain function on the type `Rectangle`. There's also a second implementation for `location()`. 

By implementing both these methods, `Rectangle` have now fully implemented the `Shape` interface.

### Pass an interface

Ok, so we've fully implemented an interface, what does it allow me to do? Well there are two things you can do:

- **Call properties and behavior**. At this point, you are ready to create an instance and call both properties and methods (its new behavior):

   ```go
   var rectangle Rectangle = Rectangle{x: 5, y: 2}
   fmt.Println(rectangle.area()) // prints 10
   ```

   Great, our `Rectangle` type has both the properties `x` and `y` as well as the behavior from `Shape`.

- **Pass an interface**. Imagine you wanted to pass the behavior to a function to make it flexible:

   ```go
   func printArea(shape Shape) {
     fmt.Println(shape.area())
   }
   ```

   To make that happen, lets change slightly how we construct our `Rectangle instance`:

   ```go
   var shape Shape = Rectangle{x: 5, y: 2}
   printArea(rectangle) // prints 10
   ```

### Implement `Square`

To really see the power in what we just created, lets create another struct `Square` and have it implement `Shape`:

```go
type Square struct {
  side int
}

func (s Square) area() int {
  return s.square * s.square
}
func (s Square) location() Point {
  return Point{x: s.side, y: s.side}
}

func main() {
  var shape Shape = Rectangle{x: 5, y: 2}
  var shape2 Shape = Square{side: 5}
  printArea(shape) // prints 10
  printArea(shape2) // prints 25
}
```

The power lies in the fact that `printArea()` doesn't have to deal with the internals of `Rectangle` or `Shape`, it just need the parameter to implement `Shape`, a behavior.

### Full code

Here's the full code:

```go
package main

import "fmt"

type Rectangle struct {
 x int
 y int
}

type Point struct {
 x int
 y int
}

type Square struct {
 side int
}

type Shape interface {
 area() int
 location() Point
}

func printArea(shape Shape) {
 fmt.Println(shape.area())
}

func (r Rectangle) area() int {
 return r.x * r.y
}

func (r Rectangle) location() Point {
 return Point{x: r.x, y: r.y}
}

func (s Square) area() int {
 return s.side * s.side
}

func (s Square) location() Point {
 return Point{x: s.side, y: s.side}
}

func main() {
 var shape Shape = Rectangle{x: 5, y: 2}
 var shape2 Shape = Square{side: 5}
 printArea(shape)  // prints 10
 printArea(shape2) // prints 25
}
```

## Type assertions

So far, a `Rectangle` or `Square` implements the `Shape` interface

Lets have a closer look at this code:

```go
var shape Shape = Rectangle{x: 5, y: 2}
var shape2 Shape = Square{side: 5}
printArea(shape)  // prints 10
printArea(shape2) // prints 25
```

We've said for `shape` and `shape2` to be of type `Shape`. That's great for being sent to the `printArea()` method. What if we need to access a `Rectangle` property on `shape`, can we? Let's try:

```go
var shape Shape = Rectangle{x: 5, y: 2}
fmt.Println(shape.x) // shape.x undefined (type Shape has no field or method x)
```

Ok, not working, we need to find a way to reach the underlying fields. We can use something called *type assertion* like so:

```go
var shape Shape = Rectangle{x: 5, y: 2}
fmt.Println(shape.(Rectangle).x) // 5
```

Ok, that works, so `.(<type>)` works, if the underlying type is the correct type. 

## Change a value

So one thing about our approach so far is that we have implemented interfaces with methods that reads data from the underlying struct instances. What if we want to change data, can we do that?

Let's look at an example:

```go
package main
import "fmt"

type Car struct {
 speed int
 model string
 make  string
}

type Runnable interface {
 run()
}

func (c Car) run() {
 c.speed = 10
}

func main() {
  c := Car{make: "Ferrari", model: "F40", speed: 0}
  c.run()
  fmt.Println(c.speed) // ?
}
```

Running this code, it returns `0`. So looking at our `run()` method:

```go
func (c Car) run() {
 c.speed = 10
}
```

shouldn't this work? Well no, because, you are not really changing the instance. For that, you need to send a reference. 

A slight alteration to the `run()` method, with `*`:

```go
func (c *Car) run() {
 c.speed = 10
}
```

and your code now does what it's supposed to.

## Assignment

Start with the following code:

```go
package main 

type Point struct {
 x float32
 y float32
}

type Vehicle struct {
 velocity float32
 Point
}

func main() {
 v := Vehicle{
  velocity: 0,
  Point: Point{
   x: 0,
   y: 0,
  },
 }
 v.fly()
 fmt.Println(v.velocity)
 v.land()
 fmt.Println(v.velocity)
}
```

Implement the following interface:

```go
type Spaceship interface {
 fly()
 land()
 position() Point
}
```

The output from running the program should be:

```output
10
0
```

## Solution

```go
package main

import "fmt"

type Point struct {
 x float32
 y float32
}

type Vehicle struct {
 velocity float32
 Point
}

type Spaceship interface {
 fly()
 land()
 position() Point
}

func (v *Vehicle) fly() {
 v.velocity = 10
}

func (v *Vehicle) land() {
 v.velocity = 0
}

func (v Vehicle) position() Point {
 return v.Point
}

func main() {
 v := Vehicle{
  velocity: 0,
  Point: Point{
   x: 0,
   y: 0,
  },
 }
 v.fly()
 fmt.Println(v.velocity)
 v.land()
 fmt.Println(v.velocity)
}

```

## Challenge

TODO

## Post-Lecture Quiz

TODO
