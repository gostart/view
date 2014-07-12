package view

type URLFunc func(ctx *Context) string

func (urlFunc URLFunc) GetURL(ctx *Context) string {
	return urlFunc(ctx)
}
