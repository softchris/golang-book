package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "   114  "

	fmt.Printf("%s , string length %d \n", s, len(s))
	res := strings.Trim(s, " ")
	fmt.Printf("%s , string length %d \n", res, len(res))

	s2 := "   114  "
	fmt.Printf("%s , string length %d \n", s2, len(s2))
	res = strings.TrimLeft(s2, " ")
	fmt.Printf("%s , string length %d \n", res, len(res))

	s3 := "   114  "
	fmt.Printf("%s , string length %d \n", s3, len(s3))
	res = strings.TrimRight(s3, " ")
	fmt.Printf("%s , string length %d \n", res, len(res))
}
