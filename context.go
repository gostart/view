package view

import (
	"fmt"
	"net/http"
)

func NewContext(server *Server, responseWriter http.ResponseWriter, httpRequest *http.Request, urlArgs ...string) *Context {
	if len(urlArgs) == 0 {
		// todo extract from httpRequest
	}
	ctx := &Context{
		Server:   server,
		Request:  newRequest(server, httpRequest),
		Response: newResponse(server, responseWriter),
		URLArgs:  urlArgs,
	}
	if server.TrackSessions {
		ctx.Session = newSession(ctx)
	}
	return ctx
}

type Context struct {
	Server   *Server
	Request  *Request
	Response *Response
	Session  *Session

	// Arguments parsed from the URL path
	URLArgs []string

	// Custom response wide data that can be set by the application
	Data      interface{}
	DebugData interface{}
}

/*
ForURLArgs returns an altered Context copy where
Context.URLArgs is set to urlArgs.
Can be used for calling the the URL() method of a URL interface
to get the URL of another view, defined by urlArgs.

The following example gets the URL of MyPage with the first
URL argument is that of the current page and the second
URL argument is "second-arg":

	MyPage.URL(ctx.ForURLArgs(ctx.URLArgs[0], "second-arg"))
*/
func (self *Context) ForURLArgs(urlArgs ...string) *Context {
	clone := *self
	clone.URLArgs = urlArgs
	return &clone
}

func (self *Context) ForURLArgsConvert(urlArgs ...interface{}) *Context {
	stringArgs := make([]string, len(urlArgs))
	for i := range urlArgs {
		stringArgs[i] = fmt.Sprint(urlArgs[i])
	}
	return self.ForURLArgs(stringArgs...)
}
