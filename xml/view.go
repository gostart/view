package xml

import (
	"encoding/xml"

	"github.com/gostart/view"
)

type View struct {
	Content interface{}
	Header  bool
	Indent  string
}

func NewView(content interface{}) *View {
	return &View{Content: content}
}

func NewViewIndent(content interface{}, indent string) *View {
	return &View{Content: content, Indent: indent}
}

func (v *View) Render(ctx *view.Context) (err error) {
	indent := v.Indent
	if indent == "" {
		indent = Config.Indent
	}
	encoder := xml.NewEncoder(ctx.Response)
	encoder.Indent("", indent)
	if v.Header {
		ctx.Response.Out(xml.Header)
	}
	return encoder.Encode(v.Content)
}
