package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type file struct {
	Name     string
	Path     string
	IsDir    bool
	SubFiles []file
}

func Main() {

	path, er := os.Getwd()
	if er != nil {
		log.Println(er)
	}

	root := file{Name: filepath.Base(path), Path: path, IsDir: true, SubFiles: nil}
	files := getFiles(path)
	root.SubFiles = files

	print(root, 1)
}

func getFiles(path string) []file {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	result := []file{}

	for _, f := range files {
		filePath := filepath.Join(path, f.Name())
		var subFiles []file
		if f.IsDir() {
			subFiles = getFiles(filePath)
		} else {
			subFiles = nil
		}
		result = append(result, file{Name: f.Name(), Path: filePath, IsDir: f.IsDir(), SubFiles: subFiles})
	}

	return result
}

func print(f file, level int) {
	prefix := ""

	for i := 1; i <= level; i++ {
		if i == level {
			prefix = prefix + "  |- "
		} else {
			prefix = prefix + "  ."
		}
	}

	fmt.Println(prefix + f.Name)
	for _, sf := range f.SubFiles {
		print(sf, level+1)
	}
}
