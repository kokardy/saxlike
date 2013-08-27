package saxlike

import (
	"bytes"
	"encoding/xml"
	"log"
	"testing"
)

type PartialHandler struct {
	VoidHandler
}

func (h PartialHandler) StartElement(element xml.StartElement) {
	log.Println("start:", element.Name.Local)
}

func TestPartialHandler(t *testing.T) {
	source := `<html> <title>taitoru</title> <body>&lt;bodfdaDF</body> </html>`
	r := bytes.NewReader([]byte(source))
	handler := &PartialHandler{}
	parser := NewParser(r, handler)
	parser.SetHTMLMode()
	parser.Parse()

}
