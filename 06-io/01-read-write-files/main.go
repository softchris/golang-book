package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	var path = "invoices.csv"
	filebuffer, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	var inputdata string = string(filebuffer)

	rows := strings.Split(inputdata, "\n")
	for _, row := range rows {
		fmt.Println("row:", row)
	}
}
