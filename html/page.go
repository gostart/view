package html

import (
	"net/http"

	"github.com/gostart/view"
)

type Page struct {
	Content view.View
}

func (page *Page) Render(ctx *view.Context) (err error) {
	return nil
}

func (page *Page) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	view.HTTPHandler(page).ServeHTTP(responseWriter, request)
}
