package main

import (
	"fmt"
)

func produceResults(ch chan int) {
	// time.Sleep(2 * time.Second)
	ch <- 1
	ch <- 2
	// ch <- 3
	close(ch)
}

func main() {
	ch := make(chan int)
	go produceResults(ch)
	// time.Sleep(1 * time.Second)

	// 	for i := 0; i < 3; i++ {
	// 		select {
	// 		case x, ok := <-ch:
	// 			if ok {
	// 				fmt.Println(x)
	// 			} else {
	// 				fmt.Println("channel closed")
	// 			}
	// 		}
	// 	}
	// }

label:
	for {
		select {
		case x, ok := <-ch:
			if ok {
				fmt.Println(x)
			} else {
				fmt.Println("channel closed")
				break label
			}
			// default:
			// fmt.Println("nothing")
		}
	}
}
