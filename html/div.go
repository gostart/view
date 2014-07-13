package html

import (
	"github.com/gostart/view"
)

// DIV creates <div class="class">content</div>
func DIV(class string, content ...interface{}) *Div {
	return &Div{Class: class, Content: view.AsView(content...)}
}

// Div represents a HTML div element.
type Div struct {
	ID      string
	Class   string
	Style   string
	OnClick string
	Content view.View
}

func (div *Div) Render(ctx *view.Context) (err error) {
	ctx.Response.Out("<div")
	if div.ID != "" {
		WriteAttrib(ctx.Response, "id", div.ID)
	}
	if div.Class != "" {
		WriteAttrib(ctx.Response, "class", div.Class)
	}
	if div.Style != "" {
		WriteAttrib(ctx.Response, "style", div.Style)
	}
	if div.OnClick != "" {
		WriteAttrib(ctx.Response, "onclick", div.OnClick)
	}
	ctx.Response.Out(">")
	if div.Content != nil {
		err = div.Content.Render(ctx)
		if err != nil {
			return err
		}
	}
	ctx.Response.Out("</div>")
	return nil
}

func (div *Div) GetID() string {
	return div.ID
}

func (div *Div) SetID(id string) {
	div.ID = id
}
