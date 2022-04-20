package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func add(first int, second int) int {
	return first + second
}

func main() {
	no := os.Args[1]
	fmt.Println(reflect.TypeOf(no))
	no1, _ := strconv.Atoi(os.Args[1])
	no2, _ := strconv.Atoi(os.Args[2])
	var sum = add(no1, no2)
	fmt.Println(sum)
	fmt.Printf("%T", os.Args[1])
	fmt.Printf("%T", 1)
}
