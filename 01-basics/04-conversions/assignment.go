package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(no int, secondNumber int) int {
	return no + secondNumber
}

func main() {
	no1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Errorf("Error converting %s to a number", os.Args[1])
		panic(err)
	}
	no2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Errorf("Error converting %s to a number", os.Args[2])
		panic(err)
	}

	var sum = add(no1, no2)
	fmt.Println(sum)
}
