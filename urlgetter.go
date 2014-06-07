package view

// URLGetter is an interface to return URL strings depending on the request path args.
type URLGetter interface {
	URL(ctx *Context) string
}
