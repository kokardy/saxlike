package saxlike

import (
	"encoding/xml"
	"io"
	"log"
)

type Parser struct {
	*xml.Decoder
	handler Handler
}

func NewParser(reader io.Reader, handler Handler) *Parser {
	decoder := xml.NewDecoder(reader)
	return &Parser{decoder, handler}
}

func (p *Parser) SetHTMLMode() {
	p.Strict = false
	p.AutoClose = xml.HTMLAutoClose
	p.Entity = xml.HTMLEntity
}

func (p *Parser) Parse() (err error) {
	p.handler.StartDocument()

	for {
		token, err := p.Token()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			panic(err)
		}
		switch token.(type) {
		case xml.StartElement:
			log.Println("start")
			s := token.(xml.StartElement)
			p.handler.StartElement(s)
		case xml.EndElement:
			log.Println("end")
			e := token.(xml.EndElement)
			p.handler.EndElement(e)
		case xml.CharData:
			log.Println("char")
			c := token.(xml.CharData)
			p.handler.CharData(c)
		case xml.Comment:
			log.Println("comment")
			com := token.(xml.Comment)
			p.handler.Comment(com)
		case xml.ProcInst:
			log.Println("proc")
			pro := token.(xml.ProcInst)
			p.handler.ProcInst(pro)
		case xml.Directive:
			log.Println("directive")
			dir := token.(xml.Directive)
			p.handler.Directive(dir)
		default:
			panic("unknown xml token.")
		}
	}

	p.handler.EndDocument()
	return
}

func Parse(decoder *xml.Decoder, handler Handler, htmlMode bool) error {
	parser := &Parser{decoder, handler}
	if htmlMode {
		parser.SetHTMLMode()
	}
	return parser.Parse()
}
