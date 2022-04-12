package main

import (
	"io/ioutil"
	"log"
	"os"
)

func ProcessFile(path string) {
	filebuffer, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	inputdata := string(filebuffer)
	log.Print("Do something with input: \n", inputdata)
}

func main() {
	fileName := "record.csv"
	logFile := "logfile"

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Could not log to file: ", logFile)
	}
	defer f.Close()

	log.SetOutput(f)

	log.Printf("processing file '%s' \n", fileName)
	ProcessFile(fileName)
}
