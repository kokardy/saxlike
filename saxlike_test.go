package saxlike

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"log"
	"strings"
	"testing"
)

func indent(i int) string {
	buf := bytes.NewBufferString("")
	for n := 0; n < i; n++ {
		buf.WriteString("\t")
	}
	return buf.String()
}

type PartialHandler struct {
	VoidHandler
	indent int
}

func (h *PartialHandler) StartElement(e xml.StartElement) {
	fmt.Printf("%s<%s>\n", indent(h.indent), e.Name.Local)
	h.indent++
}

func (h *PartialHandler) EndElement(e xml.EndElement) {
	h.indent--
	fmt.Printf("%s</%s>\n", indent(h.indent), e.Name.Local)
}

func (h *PartialHandler) CharData(c xml.CharData) {
	i := indent(h.indent)
	comment := template.HTMLEscapeString(strings.Trim(string(c), " \t\r\n"))
	if len(comment) > 0 {
		fmt.Printf("%s%s\n", i, comment)
	}
}

func (h *PartialHandler) Comment(c xml.Comment) {
	fmt.Printf("%s<!--%s-->\n", indent(h.indent), string(c))
}

func (h *PartialHandler) ProcInst(p xml.ProcInst) {
	fmt.Printf("%s<?%s %s?>\n", indent(h.indent), p.Target, p.Inst)
}

func TestPartialHandler(t *testing.T) {
	log.Println("handler test start")
	source := `<html> 
  <title>taitoru</title> 
  <body>&lt;bodfdaDF</body>
  <div>
  <!-- comment  -->
  <?target inst ?> 
  </div>
  <![CDATA[
  <fdsajk><>?fdskfdsafdsajkfldsafjkl
    fjdklasfjdkslaf
  ]]>
  </html>`
	fmt.Println(source)
	r := bytes.NewReader([]byte(source))
	handler := &PartialHandler{}
	parser := NewParser(r, handler)
	parser.SetHTMLMode()
	parser.Parse()

}
