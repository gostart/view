package view

// URLView renders an URL.
type URLView struct {
	URL
}

func (urlView URLView) Render(ctx *Context) (err error) {
	_, err = ctx.Response.WriteString(urlView.GetURL(ctx))
	return err
}
