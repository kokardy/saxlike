saxlike
=======

SAX(Simple API for XML)-like API for golang

This is just alpha version.

Handler implements:
```go
type MyHandler struct{
  depth int
}

func (h *MyHandler) StartDocument(){
  fmt.Println("doc start")
}

func (h *MyHandler) EndDocument(){
  fmt.Println("doc end")
}

func (h *MyHandler) StartElement(element xml.StartElement){
  fmt.Println("start ",element.Name.Space, ":" , element.Name.Local)
}

func (h *MyHandler) EndElement(element xml.EndElement){
  fmt.Println("end", element.Name.Space, ":" , element.Name.Local)
}

func (h *MyHandler) Comment(comment xml.Comment){
  fmt.Println("comment", string(comment))
}

func (h *MyHandler) CharData(char xml.CharData){
  fmt.Println("chardata", string(chardata))
}

func (h *MyHandler) ProcInst(proc xml.ProcInst){
  fmt.Println("proc", proc)
}

func (h *MyHandler) Directive(dir xml.Directive){
  fmt.Println("directive", string(dir))
}
```

Parse:

```go
source := `<html> <title>taitoru</title> <body>&lt;bodfdaDF</body> </html>`
r := bytes.NewReader([]byte(source))
handler := &MyHandler{}
parser := saxlike.NewParser(r, handler)
parser.SetHTMLMode()
parser.Parse()

```
