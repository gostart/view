package view

import "strings"

const URLArgPlaceholder = "(*)"

// URLString implements the URL interface for a URL string.
type URLString string

func (self URLString) GetURL(ctx *Context) string {
	url := string(self)
	for _, arg := range ctx.URLArgs {
		url = strings.Replace(url, URLArgPlaceholder, arg, 1)
	}
	return ctx.Request.AddProtocolAndHostToURL(url)
}
