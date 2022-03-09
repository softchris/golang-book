package main

import (
	"fmt"
)

type Runnable interface {
	run()
}

type Describable interface {
	description() string
}

type Car struct {
	speed int
	model string
	make  string
}

type Hero struct {
	name string
}

func DescribeThings(describable Describable) {
	fmt.Println(describable.description())
}

func RunThings(car *Car) {
	car.run()
}

func (c *Car) run() {
	c.speed = 10
}

func (c Car) description() string {
	return fmt.Sprintf("Model: %s, Make: %s", c.model, c.make)
}

func (hero Hero) description() string {
	return fmt.Sprintf("The heroes name is: %s", hero.name)
}

func main() {

	var describable Describable
	describable = Car{speed: 0, make: "Ferrari", model: "F40"}
	DescribeThings(describable)

	var describable2 Describable
	describable2 = Hero{name: "Conan"}
	DescribeThings(describable2)

	c := Car{speed: 0, make: "Porsche", model: "911"}
	p := &c
	RunThings(p)
	fmt.Println(c.speed)

}
