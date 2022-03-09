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
	fmt.Println(shape.(Rectangle).x)
}
