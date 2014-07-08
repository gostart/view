package html

import (
	"github.com/gostart/view"
)

///////////////////////////////////////////////////////////////////////////////
// Span

// Span represents a HTML span element.
type Span struct {
	ID      string
	Class   string
	Content view.View
}

func (self *Span) Render(ctx *view.Context) (err error) {
	ctx.Response.XML.OpenTag("span")
	ctx.Response.XML.AttribIfNotDefault("id", self.ID)
	ctx.Response.XML.AttribIfNotDefault("class", self.Class)
	if self.Content != nil {
		err = self.Content.Render(ctx)
	}
	ctx.Response.XML.CloseTagAlways()
	return err
}
