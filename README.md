saxlike
=======

SAX(Simple API for XML)-like API for golang

Handler implements:
```go
type CompleteHandler struct{}

func (h *CompleteHandler) StartDocument(){
  fmt.Println("doc start")
}

func (h *CompleteHandler) EndDocument(){
  fmt.Println("doc end")
}

func (h *CompleteHandler) StartElement(element xml.StartElement){
  fmt.Println("start ",element.Name.Space, ":" , element.Name.Local)
}

func (h *CompleteHandler) EndElement(element xml.EndElement){
  fmt.Println("end", element.Name.Space, ":" , element.Name.Local)
}

func (h *CompleteHandler) Comment(comment xml.Comment){
  fmt.Println("comment", string(comment))
}

func (h *CompleteHandler) CharData(char xml.CharData){
  fmt.Println("chardata", string(chardata))
}

func (h *CompleteHandler) ProcInst(proc xml.ProcInst){
  fmt.Println("proc", proc)
}

func (h *CompleteHandler) Directive(dir xml.Directive){
  fmt.Println("directive", string(dir))
}

//VoidHandler is a implemented Handler that do nothing.
type PartialHandler stuct{
  saxlike.VoidHandler
}

//You need not implement all methods.
func (h *PartialHandler) StartElement(element xml.StartElement){
  fmt.Println("start", element.Name.Space, ":", element.Name.Local)
}

```

Parse:

```go
source := `<html> <title>taitoru</title> <body>&lt;bodfdaDF</body> </html>`
r := bytes.NewReader([]byte(source))
handler := &ParticalHandler{}
parser := saxlike.NewParser(r, handler)
parser.SetHTMLMode()
parser.Parse()

```
