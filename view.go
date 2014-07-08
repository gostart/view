package view

import (
	"net/http"
)

// View is the basic interface for all types in the view package.
type View interface {
	Render(ctx *Context) (err error)
}

type ViewFunc func(ctx *Context) error

func (viewFunc ViewFunc) Render(ctx *Context) error {
	return viewFunc(ctx)
}

func HTTPHandler(view View, urlArgs ...string) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		ctx := NewContext(responseWriter, request, urlArgs...)
		err := view.Render(ctx)
		if err != nil {
			errView, ok := err.(View)
			if !ok {
				errView = ErrInternalServerError500(err)
			}
			errView.Render(ctx)
		}
	})
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

// ViewOrError returns view if err is nil, or else an Error view for err.
func ViewOrError(view View, err error) View {
	if err != nil {
		return ErrInternalServerError500(err)
	} else {
		return view
	}
}
