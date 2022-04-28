package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	var url string
	fmt.Println("Type URL: ")
	fmt.Scan(&url)
	//
	// "^([a-z]*\\.)+[a-z]*@
	r, err := regexp.Compile(`^(?P<protocol>\w+):\/\/(?P<domain>\w+\.\w+)\/(?P<route>\w+)\/?`)
	if err != nil {
		log.Fatal("Error compiling: ", err)
	}
	m := r.FindStringSubmatch(url)
	if m == nil {
		panic("mo match")
	}
	result := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = m[i]
		}
	}
	fmt.Println("The URL consist of:")
	fmt.Println(result["protocol"])
	fmt.Println(result["domain"])
	fmt.Println(result["route"])
}
