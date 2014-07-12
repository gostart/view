package html

import (
	"github.com/gostart/view"
)

// Span represents a HTML span element.
type Span struct {
	ID      string
	Class   string
	Style   string
	OnClick string
	Content view.View
}

func (span *Span) Render(ctx *view.Context) (err error) {
	ctx.Response.Out("<span")
	if span.ID != "" {
		writeAttrib(ctx.Response, "id", span.ID)
	}
	if span.Class != "" {
		writeAttrib(ctx.Response, "class", span.Class)
	}
	if span.Style != "" {
		writeAttrib(ctx.Response, "style", span.Style)
	}
	if span.OnClick != "" {
		writeAttrib(ctx.Response, "onclick", span.OnClick)
	}
	ctx.Response.Out(">")
	if span.Content != nil {
		err = span.Content.Render(ctx)
		if err != nil {
			return err
		}
	}
	ctx.Response.Out("</span>")
	return nil
}

func (span *Span) GetID() string {
	return span.ID
}

func (span *Span) SetID(id string) {
	span.ID = id
}

// SPAN creates <span class="class">content</span>
func SPAN(class string, content ...interface{}) *Span {
	return &Span{Class: class, Content: view.AsView(content...)}
}
