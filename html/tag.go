package html

import (
	"github.com/gostart/view"
)

// Tag represents an arbitrary HTML element.
type Tag struct {
	Tag     string
	ID      string
	Class   string
	Attribs map[string]string
	Content view.View
}

func (self *Tag) Render(ctx *view.Context) (err error) {
	ctx.Response.XML.OpenTag(self.Tag)
	ctx.Response.XML.AttribIfNotDefault("id", self.ID)
	ctx.Response.XML.AttribIfNotDefault("class", self.Class)
	for key, value := range self.Attribs {
		ctx.Response.XML.Attrib(key, value)
	}
	if self.Content != nil {
		err = self.Content.Render(ctx)
	}
	ctx.Response.XML.CloseTagAlways()
	return err
}
