package view

//
// See also GetViewFunc
type ViewFunc func(ctx *Context) error

func (viewFunc ViewFunc) Render(ctx *Context) error {
	if viewFunc == nil {
		return nil
	}
	return viewFunc(ctx)
}
