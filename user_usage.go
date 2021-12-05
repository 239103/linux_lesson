package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func math() {
	var result, i, n, m int64
	for i = 1; i < 100000; i++ {
		for n = 1; n < 100000; n++ {
			for m = 1; m < 100000; m++ {
				result += m * n * i
			}
		}
	}
	fmt.Println("result=", result)
}

func walkindir() {
	var files []string

	root := "/usr"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		input, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			continue
		}

		tmpfile := "/tmp" + file
		tmppath := filepath.Dir(tmpfile)
		if _, err := os.Stat(tmppath); errors.Is(err, os.ErrNotExist) {
			err := os.MkdirAll(tmppath, os.ModePerm)
			if err != nil {
				log.Println(err)
			}
		}
		err = ioutil.WriteFile(tmpfile, input, 0644)
		if err != nil {
			fmt.Println("Error creating", tmpfile)
			fmt.Println(err)
			continue
		}
	}
}

func main() {
	math()
	// walkindir()
}
