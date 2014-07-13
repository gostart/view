package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/gostart/view"
)

var CloseTag interface{}

func Element(tag string, args ...interface{}) view.View {
	lenArgs := len(args)

	var buf bytes.Buffer
	buf.WriteByte('<')
	buf.WriteString(tag)
	for i := 0; i < lenArgs; i += 2 {
		WriteAttrib(&buf, args[i], args[i+1])
	}
	if lenArgs&1 == 0 {
		buf.WriteString("/>")
		return view.String(buf.String())
	}
	buf.WriteByte('>')

	content := view.AsView(args[lenArgs-1])

	if content == nil {
		buf.WriteString("</")
		buf.WriteString(tag)
		buf.WriteByte('>')
		return view.String(buf.String())
	}

	return &elem{
		open:  buf.String(),
		view:  content,
		close: "</" + tag + ">",
	}
}

func WriteAttrib(writer io.Writer, name, value interface{}) {
	fmt.Fprint(writer, " ", name, "='")
	xml.EscapeText(writer, []byte(fmt.Sprint(value)))
	writer.Write([]byte{'\''})
}

type elem struct {
	open  string
	view  view.View
	close string
}

func (e *elem) Render(ctx *view.Context) (err error) {
	ctx.Response.Out(e.open)
	err = e.view.Render(ctx)
	if err != nil {
		return err
	}
	ctx.Response.Out(e.close)
	return nil
}
