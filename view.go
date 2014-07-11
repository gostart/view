package view

// View is the basic interface for all types in the view package.
type View interface {
	Render(ctx *Context) (err error)
}

func NewView(val interface{}) View {
	if val == nil {
		return nil
	}
	switch s := val.(type) {
	case View:
		return s
	case *View:
		return ViewPtr{s}
	case string:
		return String(s)
	}
	return &Print{val}
}

type ViewPtr struct {
	Ptr *View
}

func (viewPtr ViewPtr) Render(ctx *Context) error {
	if viewPtr.Ptr == nil || *viewPtr.Ptr == nil {
		return nil
	}
	return (*viewPtr.Ptr).Render(ctx)
}

type ViewFunc func(ctx *Context) error

func (viewFunc ViewFunc) Render(ctx *Context) error {
	if viewFunc == nil {
		return nil
	}
	return viewFunc(ctx)
}
