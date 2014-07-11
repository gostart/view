package view

import (
	"net/http"
)

func HTTPHandler(server *Server, view View, urlArgs ...string) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		ctx := NewContext(server, responseWriter, request, urlArgs...)
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

// ProductionServerView returns view if server.IsProductionServer
// is true, else nil which is a valid value for a View.
func ProductionServerView(server *Server, view View) View {
	if !server.IsProductionServer {
		return nil
	}
	return view
}

// NonProductionServerView returns view if server.IsProductionServer
// is false, else nil which is a valid value for a View.
func NonProductionServerView(server *Server, view View) View {
	if server.IsProductionServer {
		return nil
	}
	return view
}

// ViewOrError returns view if err is nil, or else an Error view for err.
func ViewOrError(view View, err error) View {
	if err == nil {
		return view
	} else {
		return ErrInternalServerError500(err)
	}
}
