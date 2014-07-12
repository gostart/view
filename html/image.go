package html

import (
	"github.com/gostart/view"
)

type Image struct {
	view.URL
	ID      string
	Class   string
	Width   int
	Height  int
	Title   string
	OnClick string
}

func (img *Image) Render(ctx *view.Context) (err error) {
	ctx.Response.Out("<img")
	if img.ID != "" {
		writeAttrib(ctx.Response, "id", img.ID)
	}
	if img.Class != "" {
		writeAttrib(ctx.Response, "class", img.Class)
	}
	if img.Width != 0 {
		writeAttrib(ctx.Response, "width", img.Width)
	}
	if img.Height != 0 {
		writeAttrib(ctx.Response, "height", img.Height)
	}
	if img.Title != "" {
		writeAttrib(ctx.Response, "title", img.Title)
		writeAttrib(ctx.Response, "alt", img.Title)
	}
	if img.OnClick != "" {
		writeAttrib(ctx.Response, "onclick", img.OnClick)
	}
	writeAttrib(ctx.Response, "src", img.GetURL(ctx))
	ctx.Response.Out("/>")
	return nil
}

func (img *Image) GetID() string {
	return img.ID
}

func (img *Image) SetID(id string) {
	img.ID = id
}

// IMG creates a HTML img element for an URL with optional width and height.
// The first int of dimensions is width, the second one height.
func IMG(url string, dimensions ...int) *Image {
	width := 0
	height := 0
	dimCount := len(dimensions)
	if dimCount >= 1 {
		width = dimensions[0]
		if dimCount >= 2 {
			height = dimensions[1]
		}
	}
	return &Image{URL: view.URLString(url), Width: width, Height: height}
}
