package dir

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

var ErrDirExist = errors.New("Dir exist error")

/*
DESCRIPTION

Reads the content on a directory and returns `FileInfo` array and a possible error, if dir can't be read.

PARAMS

path:string, a string representing the path holding a directory

EXAMPLE

	files, err := dir.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}

*/
func ReadDir(path string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func CreateDir(dirName string) error {
	_, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			return err
		}
		return nil
	} else if err != nil {
		fmt.Println("error creating dir")
		return err // other type of error
	} else {
		return ErrDirExist
	}
}
