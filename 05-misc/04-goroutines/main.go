package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

func SearchFiles(dir string, lookFor string, ch chan string) {
	log.Println("[SEARCHING] ", dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		log.Println(dir+file.Name(), file.IsDir())
		if file.Name() == lookFor {
			ch <- "[FOUND] " + filepath.Join(dir, file.Name())
			return
		}
	}
	ch <- "[NOT FOUND] " + dir
}

func main() {
	ch := make(chan string)

	go SearchFiles("./test/", "test2.txt", ch)
	go SearchFiles("./other/", "test2.txt", ch)

	var res = ""
	for i := 0; i < 2; i++ {
		res = <-ch
		log.Println(res)
	}
}
