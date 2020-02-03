package main

import (
	"fmt"
	"io/ioutil"
	"main/DocStract"
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

	for a, binaryData := range *files {
		err := ioutil.WriteFile(fmt.Sprintf("%d.pdf", a), binaryData.Bytes, 0644)
		if err != nil {
			panic(err)
		}
	}
}
