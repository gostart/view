package html

import (
	"github.com/gostart/view"
)

// DIV creates <div class="class">content</div>
func DIV(class string, content ...interface{}) *Div {
	return &Div{Class: class, Content: WrapContents(content...)}
}

// Div represents a HTML div element.
type Div struct {
	ID      string
	Class   string
	Style   string
	Content view.View
	OnClick string
}

func (self *Div) Render(ctx *view.Context) (err error) {
	ctx.Response.XML.OpenTag("div")
	ctx.Response.XML.AttribIfNotDefault("id", self.ID)
	ctx.Response.XML.AttribIfNotDefault("class", self.Class)
	ctx.Response.XML.AttribIfNotDefault("style", self.Style)
	ctx.Response.XML.AttribIfNotDefault("onclick", self.OnClick)
	if self.Content != nil {
		err = self.Content.Render(ctx)
	}
	ctx.Response.XML.CloseTagAlways()
	return err
}
