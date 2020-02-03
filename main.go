package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/richardlehane/mscfb"
)

func main() {
	data, err := ioutil.ReadFile("2017-05-17_NewDelhiExpress_Off-Hire_HLAG.msg")

	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(data)

	doc, err := mscfb.New(reader)

	if err != nil {
		panic(err)
	}

	files := [][]byte{}
	attachment := false
	file := 0

	for entry, err := doc.Next(); err == nil; entry, err = doc.Next() {

		//fmt.Println(entry.Name, entry.Size, entry.Path, entry.Initial)

		if strings.Contains(entry.Name, "attach") {
			files = append(files, []byte{})
			attachment = true
			continue
		}
		if attachment && strings.Contains(entry.Name, "properties") {
			attachment = false
			file++
			continue
		}

		if attachment {
			buf := make([]byte, entry.Size)
			i, _ := entry.Read(buf)
			if i > 0 {
				files[file] = append(files[file], buf[:i]...)
			}
		}

	}

	fmt.Println("Found ", len(files), " files what can we do with these?")

	for a, binaryData := range files {
		err := ioutil.WriteFile(fmt.Sprintf("%d.pdf", a), binaryData, 0644)
		if err != nil {
			panic(err)
		}
	}
}
