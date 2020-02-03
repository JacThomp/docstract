package DocStract

import (
	"strings"

	"github.com/sirupsen/logrus"
)

//DocType is a wrapper for the type iota/enum
type DocType int

const (
	//DocUnkown represents an unknown document type
	DocUnkown = iota

	//DocPDF represents a pdf document type
	DocPDF

	//DocX represents a microsoft docx document type
	DocX
)

//DocStract stores the binary data for extracted files, as well as the type and filename metadata
type DocStract struct {
	Type     DocType
	FileName *string
	Bytes    []byte
}

//sets name to nil if cannot dertermine name and type to unkown
func (d *DocStract) getName() {
	blocks := strings.Split(string(d.Bytes), "\n")
	nameBlock := blocks[len(blocks)-1]

	chunks := strings.Split(nameBlock, ".")

	nameChunk := 0
	typeChunk := 0

	switch len(chunks[0]) {
	case 0: //pdf
		nameChunk = 2
	default: //html
		nameChunk = 0
	}

	for i := nameChunk + 1; i < len(chunks); i++ {
		if len(chunks[i]) > 1 {
			typeChunk = i
			break
		}
	}

	name := strings.TrimSpace(chunks[nameChunk])
	t := strings.TrimSpace(chunks[typeChunk])
	logrus.Info(name, ":", t, "::", len(chunks))

	/*name = StripSeperators(name)
	t = StripSeperators(t)
	logrus.Info(name, ":", t)*/

	if strings.Contains(t, "pdf") {
		name += ".pdf"
		d.Type = DocPDF
	}

	d.FileName = &name
}
