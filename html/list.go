package html

import (
	"fmt"

	"github.com/gostart/view"
)

type List struct {
	ID          string
	Class       string
	Style       string
	Ordered     bool
	OrderOffset uint
	Items       []view.View
}

func (list *List) Render(ctx *view.Context) (err error) {
	if list.Ordered {
		ctx.Response.Out("<ol")
		if list.OrderOffset != 0 {
			WriteAttrib(ctx.Response, "start", list.OrderOffset+1)
		}
	} else {
		ctx.Response.Out("<ul")
	}
	if list.ID != "" {
		WriteAttrib(ctx.Response, "id", list.ID)
	}
	if list.Class != "" {
		WriteAttrib(ctx.Response, "class", list.Class)
	}
	if list.Style != "" {
		WriteAttrib(ctx.Response, "style", list.Style)
	}

	for i, view := range list.Items {
		ctx.Response.Out("<li")
		if list.ID != "" {
			WriteAttrib(ctx.Response, "id", fmt.Sprint(list.ID, "_", i))
		}
		ctx.Response.Out(">")
		if view != nil {
			err = view.Render(ctx)
			if err != nil {
				return err
			}
		}
		ctx.Response.Out("</li>")
	}

	if list.Ordered {
		ctx.Response.Out("</ol>")
	} else {
		ctx.Response.Out("</ul>")
	}
	return nil
}

func (list *List) GetID() string {
	return list.ID
}

func (list *List) SetID(id string) {
	list.ID = id
}

// UL is a shortcut to create an unordered list by wrapping items as HTML views.
// NewView will be called for every passed item.
//
// Example:
//   UL("red", "green", "blue")
//   UL(A(url1, "First Link"), A(url2, "Second Link"))
//
func UL(items ...interface{}) *List {
	return &List{Items: view.AsViews(items...)}
}

// OL is a shortcut to create an ordered list by wrapping items as HTML views.
// NewView will be called for every passed item.
//
// Example:
//   OL("red", "green", "blue")
//   OL(A(url1, "First Link"), A(url2, "Second Link"))
//
func OL(items ...interface{}) *List {
	list := UL(items...)
	list.Ordered = true
	return list
}
