package saxlike

import (
	"encoding/xml"
)

//SAX-like handler
type Handler interface {
	//called when XML document start
	StartDocument()
	//called when XML document end
	EndDocument()
	//called when XML tag start
	StartElement(xml.StartElement)
	//called when XML tag end
	EndElement(xml.EndElement)
	//called when the parser encount chardata
	CharData(xml.CharData)
	//called when the parser encount comment
	Comment(xml.Comment)
	//called when the parser encount procInst
	//<!procinst >
	ProcInst(xml.ProcInst)
	//called when the parser encount directive
	//
	Directive(xml.Directive)
}

//VoidHandler is a implemented Handler
//All methods do nothing
//You need not implement all Handler methods.
/*
type PartialHandler struct{VoidHandler}
func (h PartialHandler) StartElement(xml.StartElement){
  //do something
}
*/
type VoidHandler struct{}

func (h VoidHandler) StartDocument()                {}
func (h VoidHandler) EndDocument()                  {}
func (h VoidHandler) StartElement(xml.StartElement) {}
func (h VoidHandler) EndElement(xml.EndElement)     {}
func (h VoidHandler) CharData(xml.CharData)         {}
func (h VoidHandler) Comment(xml.Comment)           {}
func (h VoidHandler) ProcInst(xml.ProcInst)         {}
func (h VoidHandler) Directive(xml.Directive)       {}
