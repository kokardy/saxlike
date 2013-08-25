package saxlike
import (
  "encoding/xml"
  "io"
)

//SAX-like handler
type Handler interface{
  StartDocument()
  EndDocument()
  StartElement(xml.StartElement)
  EndElement(xml.EndElement)
  CharData(xml.CharData)
  Comment(xml.Comment)
  ProcInst(xml.ProcInst)
  Directive(xml.Directive)
}

type Parser Struct{
  *xml.Decoder
  handler *Handler
}

func NewParser(reader io.Reader, handler Handler){
  decoder := xml.NewDecoder(reader)
  return &Parser{decoder, handler}
}

func (p *Parser)SetHTMLMode(){
  p.Strict = false
	p.AutoClose = xml.HTMLAutoClose
	p.Entity = xml.HTMLEntity
}

func (p *Parser)Parse() (err error){
  p.handler.StartDocument()
  
  for {
    token, err := p.Token()
    if err != io.EOF{
      err = nil
      break
    }
    if err != nil{
      return
    }
    switch token.(type){
      case xml.StartElement:
        s := token.(xml.StartElement)
        p.handler.StartElement(s)
      case xml.EndElement:
        e := token.(xml.EndElement)
        p.handler.EndElement(e)
      case xml.CharData:
        c := token.(xml.CharData)
        p.handler.CharData(c)
      case xml.Comment(xml.Comment)
        com := token.(xml.Comment)
        p.handler.Comment(com)
      case xml.ProcInst:
        pro := token.(xml.ProcInst)
        p.handler.ProcInst(pro)
      case xml.Directive:
        dir := token.(xml.Directive)
        p.handler.Directive(dir)
      default:
        panic("unknown xml token.")
    }
  }
  
  p.handler.EndDocument()
}

func Parse(decoder *xml.Decoder, handler Handler) error{
  parser := &Parser{decoder, handler)
  return parser.Parse()
}
