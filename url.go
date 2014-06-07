package view

import "strings"

const URLArgPlaceholder = "(*)"

// URL implements the URLGetter interface for a URL string.
type URL string

func (self URL) URL(ctx *Context) string {
	url := string(self)
	for _, arg := range ctx.URLArgs {
		url = strings.Replace(url, URLArgPlaceholder, arg, 1)
	}
	return ctx.Request.AddProtocolAndHostToURL(url)
}
