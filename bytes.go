package view

type Bytes []byte

func (bytes Bytes) Render(ctx *Context) error {
	_, err := ctx.Response.Write([]byte(bytes))
	return err
}
