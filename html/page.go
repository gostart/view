package html

import (
	// "net/http"

	"github.com/gostart/view"
)

type Page struct {
	Content view.View
}

func (page *Page) Render(ctx *view.Context) (err error) {
	return nil
}
