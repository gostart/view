package view

// GetViewFunc is a function pointer that implements View.
// At every Render method call, the function will be called
// and Render will be called at the result View if it is not nil.
//
// See also ViewFunc
type GetViewFunc func(ctx *Context) View

func (getView GetViewFunc) Render(ctx *Context) (err error) {
	if getView != nil {
		if view := getView(ctx); view != nil {
			return view.Render(ctx)
		}
	}
	return nil
}
