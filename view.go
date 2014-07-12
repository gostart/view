package view

import ()

// View is the basic interface for all types in the view package.
type View interface {
	Render(ctx *Context) (err error)
}

func asView(value interface{}) View {
	if value == nil {
		return nil
	}
	switch s := value.(type) {
	case View:
		return s

	case *View:
		return Pointer{s}

	case func(ctx *Context) error:
		return ViewFunc(s)

	case error:
		return Error{s}

	case string:
		return String(s)

	case []byte:
		return Bytes(s)

	case URL:
		return URLView{s}
	}
	return Print(value)
}

func AsView(values ...interface{}) View {
	if len(values) == 0 {
		return nil
	}
	if len(values) == 1 {
		return asView(values[0])
	}
	return AsViews(values...)
}
