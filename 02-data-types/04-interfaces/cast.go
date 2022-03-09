package main

import (
	"errors"
	"fmt"
	"log"
)

func elementAt(elements interface{}, index int) (interface{}, error) {
	switch t := elements.(type) {
	case []int:
		return t[index], nil
	case []string:
		return t[index], nil
	default:
		fmt.Println(t)
		return nil, errors.New("unsupported")
	}
}

func main() {
	intArr := []int{1, 2, 3}
	stringArr := []string{"one", "two", "three"}

	value, err := elementAt(intArr, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
	value, err = elementAt(stringArr, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
