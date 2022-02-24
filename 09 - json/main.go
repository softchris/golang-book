package main

import (
	"encoding/json"
	"fmt"
)

// ok, so the names have to be uppercase, or no dice
type Person struct {
	Name string
	Age  int
}

// with instructed names
type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("hello")
	res := &Person{
		Name: "chris",
		Age:  20,
	}
	jsonResult, _ := json.Marshal(res)
	fmt.Println(string(jsonResult))

	res2 := &Product{
		Id:   1,
		Name: "Tomato",
	}
	jsonResult2, _ := json.Marshal(res2)
	fmt.Println(string(jsonResult2))
}
