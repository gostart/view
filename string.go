package view

type String string

func (str String) Render(ctx *Context) (err error) {
	ctx.Response.Print(string(str))
	return nil
}
