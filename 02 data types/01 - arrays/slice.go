package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	slice := make([]int, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice2 := make([]int, 2, 5)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	dest := make([]int, 5)
	copy(dest, arr[0:2]) // copies slice into dest
	fmt.Println(dest)
}
