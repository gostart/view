package json

import (
	"encoding/json"

	"github.com/gostart/view"
)

type View struct {
	Content interface{}
	Indent  string
}

func NewView(content interface{}) *View {
	return &View{Content: content}
}

func NewViewIndent(content interface{}, indent string) *View {
	return &View{Content: content, Indent: indent}
}

func (v *View) Render(ctx *view.Context) (err error) {
	encoder := json.NewEncoder(ctx.Response)
	err = encoder.Encode(v.Content)
	if err != nil {
		return err
	}
	indent := v.Indent
	if indent == "" {
		indent = Config.Indent
	}
	if indent != "" {
		data := ctx.Response.GetBytesAndReset()
		err = json.Indent(&ctx.Response.Buffer, data, "", indent)
	}
	return err
}
