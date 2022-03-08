package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(first int, second int) int {
	return first + second
}

func add2(first int, second int) (sum int) {
	sum = first + second
	return
}

func add3(first string, second string) int {
	int1, _ := strconv.Atoi(first)
	int2, _ := strconv.Atoi(second)
	return int1 + int2
}

func calc(first int, second int) (sum int, product int) {
	sum = first + second
	product = first * second
	return
}

func changeName(name *string) {
	*name = "sara"
}

func main() {
	var name = "chris"
	sum := add(1, 2)
	fmt.Println(sum)
	fmt.Println(add2(2, 2))
	fmt.Println(add3(os.Args[1], os.Args[2]))
	sum, mult := calc(1, 2)
	fmt.Println(sum, mult)
	fmt.Println(name)
	changeName(&name)
	fmt.Println(name)
}
