package bootstrap

import (
	"github.com/gostart/view"
	"github.com/gostart/view/html"
)

func Jumbotron(content ...interface{}) view.View {
	return html.DIV("jumbotron", content...)
}

func PageHeader(content ...interface{}) view.View {
	return html.DIV("page-header", content...)
}
