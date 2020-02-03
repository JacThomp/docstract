package main

import (
	"fmt"
	"io/ioutil"
	"main/DocStract"

	"github.com/sirupsen/logrus"
)

func main() {
	data, err := ioutil.ReadFile("2017-05-17_NewDelhiExpress_Off-Hire_HLAG.msg")

	if err != nil {
		panic(err)
	}

	files, count, err := DocStract.Extract(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Found ", count, " files what can we do with these?")

	for _, document := range *files {
		if err := document.SaveFile(""); err != nil {
			logrus.Warn(err)
		}
	}
}
