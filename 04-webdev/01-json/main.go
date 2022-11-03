package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{ "name": "chris", "age": 20 }`
	person := Person{}
	json.Unmarshal([]byte(str), &person)
	fmt.Println(person)
}
