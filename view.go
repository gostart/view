package view

// View is the basic interface for all types in the view package.
type View interface {
	Render(ctx *Context) (err error)
}

type ViewFunc func(ctx *Context) error

func (viewFunc ViewFunc) Render(ctx *Context) error {
	return viewFunc(ctx)
}

// ProductionServerView returns view if view.Config.IsProductionServer
// is true, else nil which is a valid value for a View.
func ProductionServerView(view View) View {
	if !Config.IsProductionServer {
		return nil
	}
	return view
}

// NonProductionServerView returns view if view.Config.IsProductionServer
// is false, else nil which is a valid value for a View.
func NonProductionServerView(view View) View {
	if Config.IsProductionServer {
		return nil
	}
	return view
}
