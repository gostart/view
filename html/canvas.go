package html

import (
	"github.com/gostart/view"
)

// CANVAS creates <canvas class="class" width="width" height="height"></canvas>
func CANVAS(class string, width, height int) *Canvas {
	return &Canvas{Class: class, Width: width, Height: height}
}

type Canvas struct {
	ID     string
	Class  string
	Width  int
	Height int
}

func (canvas *Canvas) Render(ctx *view.Context) (err error) {
	ctx.Response.XML.OpenTag("label")
	ctx.Response.XML.AttribIfNotDefault("id", canvas.ID)
	ctx.Response.XML.AttribIfNotDefault("class", canvas.Class)
	ctx.Response.XML.Attrib("width", canvas.Width).Attrib("height", canvas.Height)
	ctx.Response.XML.CloseTagAlways()
	return err
}
