package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

func errorHandler() {
	if r := recover(); r != nil {
		// fmt.Println("Recovered ", r)
		// fmt.Println(string(debug.Stack()))
		log.Println(r, string(debug.Stack()))
	}
}

func Divide(nominator int, divider int) float32 {
	defer errorHandler()
	if divider == 0 {
		panic("can't divide by 0")
	}
	return float32(nominator) / float32(divider)
}

func main() {
	f, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(f)

	log.Println("starting program")
	no := Divide(10, 0)
	fmt.Println(no)

	no = Divide(10, 1)
	fmt.Println(no)
	f.Close()
}
