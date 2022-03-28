package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

var ErrorDivideBeZero = errors.New("Divide by zero")

func Divide2(nominator int, divider int) (float32, error) {
	if divider <= 0 {
		return 0, ErrorDivideBeZero
	}
	return float32(nominator) / float32(divider), nil
}

func Divide(nominator int, divider int) float32 {
	if divider <= 0 {
		panic("divider below zero")
	}
	return float32(nominator) / float32(divider)
}

func main() {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	defer f.Close()
	log.SetOutput(f)

	log.Println("starting batch job", time.Now())
	fmt.Println(Divide(10, 2))
	val, err := Divide2(10, 0)
	if err != nil {
		log.Fatalln(err)
		// fmt.Println(err)
		// os.Exit(-1)
	}
	fmt.Println(val)
	log.Println("stopping batch job", time.Now())
}
