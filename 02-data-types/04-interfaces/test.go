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

func (c *Car) run() {
	c.speed = 10
}

func main() {
	c := Car{make: "Ferrari", model: "F40", speed: 0}
	c.run()
	fmt.Println(c.speed) // ?
}
