package saxlike

import (
  "encoding/xml"
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
