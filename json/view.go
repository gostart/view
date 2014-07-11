package json

import (
	"encoding/json"
	"net/http"

	"github.com/gostart/view"
)

type View struct {
	Content interface{}
	Indent  string
}

func (v *View) Render(ctx *view.Context) (err error) {
	var data []byte
	indent := v.Indent
	if indent == "" {
		indent = Config.Indent
	}
	if indent == "" {
		data, err = json.Marshal(v.Content)
	} else {
		data, err = json.MarshalIndent(v.Content, "", Config.Indent)
	}
	if err != nil {
		return err
	}
	ctx.Response.Write(data)
	return nil
}
