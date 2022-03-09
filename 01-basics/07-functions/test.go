package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	no, err := strconv.Atoi(os.Args[1])
	fmt.Println(no)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Couldn't convert: " + os.Args[1])

	} else {
		fmt.Println(no)
	}
}
