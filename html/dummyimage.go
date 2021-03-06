package html

import (
	"fmt"
	"net/url"

	"github.com/gostart/view"
)

///////////////////////////////////////////////////////////////////////////////
// DummyImage

func NewDummyImage(width, height int) *DummyImage {
	return &DummyImage{Width: width, Height: height}
}

// DummyImage represents a HTML img element with src utilizing http://dummyimage.com.
type DummyImage struct {
	ID              string
	Class           string
	Width           int
	Height          int
	BackgroundColor string
	ForegroundColor string
	Text            string
}

func (self *DummyImage) Render(ctx *view.Context) (err error) {
	src := fmt.Sprintf("http://dummyimage.com/%dx%d", self.Width, self.Height)

	if self.BackgroundColor != "" || self.ForegroundColor != "" {
		if self.BackgroundColor != "" {
			src += "/" + self.BackgroundColor
		} else {
			src += "/ccc"
		}

		if self.ForegroundColor != "" {
			src += "/" + self.ForegroundColor
		}
	}

	src += ".png"

	if self.Text != "" {
		src += "&text=" + url.QueryEscape(self.Text)
	}

	// ctx.Response.XML.OpenTag("img")
	// ctx.Response.XML.AttribIfNotDefault("id", self.ID)
	// ctx.Response.XML.AttribIfNotDefault("class", self.Class)
	// ctx.Response.XML.Attrib("src", src)
	// ctx.Response.XML.AttribIfNotDefault("width", self.Width)
	// ctx.Response.XML.AttribIfNotDefault("height", self.Height)
	// ctx.Response.XML.AttribIfNotDefault("alt", self.Text)
	// ctx.Response.XML.CloseTag()
	return nil
}
