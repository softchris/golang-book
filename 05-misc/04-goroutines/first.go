package main

import (
	"fmt"
	"time"
)

func myFunction() {
	time.Sleep(1500 * time.Millisecond)
	for i := 0; i < 3; i++ {
		fmt.Println("my function: ", i)
	}
}
func anotherFunction() {
	time.Sleep(500 * time.Millisecond)
	for i := 4; i < 7; i++ {
		fmt.Println("another function: ", i)
	}
}

func main() {
	go myFunction()
	go anotherFunction()
	time.Sleep(2 * time.Second)
}
