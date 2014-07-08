package xml

import (
	"encoding/xml"
	"net/http"

	"github.com/gostart/view"
)

type View struct {
	Content interface{}
	Header  bool
	Indent  string
}

func (v *View) Render(ctx *view.Context) (err error) {
	var data []byte
	indent := v.Indent
	if indent == "" {
		indent = Config.Indent
	}
	if indent == "" {
		data, err = xml.Marshal(v.Content)
	} else {
		data, err = xml.MarshalIndent(v.Content, "", Config.Indent)
	}
	if err != nil {
		return err
	}
	if v.Header {
		ctx.Response.Print(xml.Header)
	}
	ctx.Response.Write(data)
	return nil
}

func (v *View) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	view.HTTPHandler(v).ServeHTTP(responseWriter, request)
}
