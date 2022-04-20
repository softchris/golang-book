package main

import (
	"fmt"
	"time"
)

func run(ch chan int, no int) {
	ch <- no
}

func main() {
	ch := make(chan int)

	go run(ch, 1)
	go run(ch, 2)
	time.Sleep(1 * time.Second)

block:
	for {
		select {
		case x, ok := <-ch:
			if ok {
				fmt.Println(x)
			} else {
				fmt.Println("channel closed")

				break block
			}
		default:
			fmt.Println("No value to read, exiting")

			break block
		}
	}
	fmt.Println("Done with values")

}
