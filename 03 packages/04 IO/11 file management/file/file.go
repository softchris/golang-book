package file

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

func OpenText(path string) (string, error) {
	filebuffer, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	var inputdata string = string(filebuffer)
	return inputdata, nil
}

func GetFileInfo(path string) (fs.FileInfo, error) {
	fileStat, err := os.Stat(path)
	return fileStat, err
}

func Append(path string, content string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "%s\n", content)
	return nil
}

func WriteText(path string, content string) error {
	f, err := os.Create(path)

	if err != nil {
		return err
	}
	n, err := f.WriteString(content + "\n")
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n)
	f.Sync()
	return nil
}

func CopyFile(src string, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	newFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}

func RenameFile(src string, dest string) error {
	err := os.Rename(src, dest)
	return err
}

func RemoveFile(path string) error {
	err := os.Remove(path)
	return err
}
