package view

import (
	"net/http"
)

type Template struct {
	Filename       string // Will set file extension at ContentType
	Text           string
	ContentTypeExt string
}

func (template *Template) Render(ctx *Context) (err error) {
	return nil
}

func (template *Template) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	HTTPHandler(template).ServeHTTP(responseWriter, request)
}
