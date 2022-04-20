package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	matched, err := regexp.MatchString("[a-z]{1}", "hi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("matched", matched)
}
