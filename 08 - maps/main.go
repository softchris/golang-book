package main

import "fmt"

func main() {
	cars := map[string]string{"make": "Ferrari", "model": "F40"}
	fmt.Println(cars["make"])
	fmt.Printf("cars\t%v\n", cars)
}
