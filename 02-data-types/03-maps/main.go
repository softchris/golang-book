package main

import "fmt"

func main() {
	cars := map[string]string{"make": "Ferrari", "model": "F40"}
	fmt.Println(cars["make"])
	fmt.Printf("cars\t%v\n", cars)
	_, exist := cars["model"]

	fmt.Println(exist)
	val, exist := cars["m"]
	fmt.Println(val)
	fmt.Println(exist)
}
