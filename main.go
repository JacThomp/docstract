package main

import (
	"fmt"
	"io/ioutil"
	"main/DocStract"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	path, _ := os.Getwd()
	pathFiles, _ := ioutil.ReadDir(path)

	for _, file := range pathFiles {
		if file.Name()[len(file.Name())-4:] != ".msg" {
			continue
		}

		data, err := ioutil.ReadFile(file.Name())

		if err != nil {
			panic(err)
		}

		files, count, err := DocStract.Extract(data)
		if err != nil {
			panic(err)
		}

		fmt.Println("Found ", count, " files")

		for _, document := range *files {
			if err := document.SaveFile(""); err != nil {
				logrus.Warn(err)
			} else {
				logrus.Infof("Saved file %s", *document.FileName)
			}
		}
	}
}
