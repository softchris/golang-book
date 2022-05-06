package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"iohelper/dir"
	"iohelper/file"
	"log"
)

type Products struct {
	Products []Product `json: products`
}

type Product struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
}

// writes JSON to file
// TODO

// reads JSON from file
func OpenJson() {
	file, _ := ioutil.ReadFile("products.json")

	data := Products{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Products); i++ {
		fmt.Println("Product Id: ", data.Products[i].Id)
		fmt.Println("Name: ", data.Products[i].Name)
	}
}

func main() {
	// open file
	content, err := file.OpenText("test.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("file content", content)
	}

	err = file.WriteText("test2.txt", "here's some content")
	if err != nil {
		log.Fatal(err)
	}

	err = file.Append("test.txt", "append this")
	if err != nil {
		log.Fatal(err)
	}

	files, err := dir.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		log.Println(file.Name())
	}

	err = dir.CreateDir("tmp")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dir created")

	file.CopyFile("test.txt", "copy.txt")
	log.Println("file copied")

	err = file.RenameFile("test.txt", "renamed.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileStat, err := file.GetFileInfo("renamed.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileStat.Name())        // Base name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", fileStat.IsDir())   // Abbreviation for Mode().IsDir()

	err = file.RenameFile("copy.txt", "test.txt")

	if err != nil {
		log.Fatal(err)
	}
	log.Println("file renamed")

	err = file.RemoveFile("renamed.txt")

	if err != nil {
		log.Fatal(err)
	}
	log.Println("file removed")

	OpenJson()

	/*
		  CHECK - Append text to file
			Create empty file
			CHECK - Read dir
			CHECK Create dir
			CHECK Copy file
		  Move file
			CHECK Get file metadata
			CHECK Delete fil
	*/
}
