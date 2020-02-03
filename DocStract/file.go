package DocStract

import (
	"errors"
	"io/ioutil"
	"strings"
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

	//DocXLSX represents microsoft excel doc
	DocXLSX

	//DocHTML represents an html document type
	DocHTML
)

//DocStract stores the binary data for extracted files, as well as the type and filename metadata
type DocStract struct {
	Type     DocType
	FileName *string
	Bytes    []byte
}

//Saves the file to the path, does not check if it's an unkown filetype only if it has a name
func (d *DocStract) SaveFile(path string) error {
	if len(path) > 0 && path[len(path)-1] != '/' {
		path += "/"
	}

	if d.FileName == nil {
		return errors.New("document does not have a filename cannot save")
	}

	return ioutil.WriteFile(path+*(d.FileName), d.Bytes, 0644)
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
		switch {
		case strings.Contains(chunks[0], "word"): //docx
			nameChunk = 8
			break
		case strings.Contains(chunks[2], "worksheets"): //xlsx
			for i := 3; i < len(chunks); i++ {
				if strings.Contains(StripSeperators(chunks[i]), "xlsx") {
					nameChunk = i + 1
					break
				}
			}
			break
		default: //html
			nameChunk = 0
		}
	}

	for i := nameChunk + 1; i < len(chunks); i++ {
		if len(chunks[i]) > 1 {
			typeChunk = i
			break
		}
	}

	name := strings.TrimSpace(chunks[nameChunk])
	t := strings.TrimSpace(chunks[typeChunk])

	name = StripSeperators(name)
	t = StripSeperators(t)

	switch {
	case strings.Contains(t, "pdf"):
		name += ".pdf"
		name = name[3:]
		d.Type = DocPDF
		break

	case strings.Contains(t, "docx"):
		name += ".docx"
		name = name[3:]
		d.Type = DocX
		break

	case strings.Contains(t, "xlsx"):
		name += ".xlsx"
		head := 0
		for i, b := range d.Bytes {
			if i+1 < len(d.Bytes) && b == 'P' && d.Bytes[i+1] == 'K' {
				head = i
				break
			}
		}
		name = name[3:]
		d.Type = DocXLSX
		d.Bytes = d.Bytes[head:]
		break

	case strings.Contains(t, "htm"):
		name += ".html"
		d.Type = DocHTML
		break
	}

	d.FileName = &name
}
