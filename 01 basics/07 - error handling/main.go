package main

import (
	"errors"
	"fmt"
)

var NoTooSmall = errors.New("the number is too small")

func ReturnPositive(no int) (int, error) {
	if no > 0 {
		return no, nil
	} else {
		return 0, NoTooSmall
	}

}

func main() {
	fmt.Println("hi")
	no, err := ReturnPositive(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(no)
	}

}
