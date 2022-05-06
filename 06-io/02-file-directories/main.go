package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func GetType(isDir bool) string {
	if isDir {
		return "directory"
	}
	return "file"
}

func main() {
	var path string = "tmp"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reading directory ", path)
	fmt.Println("Name, Type, Size, Modified")
	for _, file := range files {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("File Name: %s, ", file.Name())
		fmt.Printf("Type: %s, ", GetType(file.IsDir()))
		fmt.Printf("Size: %d, ", file.Size())             // Length in bytes for regular files
		fmt.Printf("Last Modified: %s, ", file.ModTime()) // Last modification time
		fmt.Print("\n")
	}
}
