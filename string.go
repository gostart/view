package view

type String string

func (str String) Render(ctx *Context) (err error) {
	_, err = ctx.Response.WriteString(string(str))
	return err
}
