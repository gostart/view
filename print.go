package view

type Print struct {
	Val interface{}
}

func (print *Print) Render(ctx *Context) error {
	_, err := ctx.Response.Print(print.Val)
	return err
}

// type Printf struct {
// 	Format string
// 	Args   []interface{}
// }

// func (printf *Printf) Render(ctx *Context) error {
// 	_, err := ctx.Response.Printf(printf.Format, printf.Args...)
// 	return err
// }
