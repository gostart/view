package view

///////////////////////////////////////////////////////////////////////////////
// SessionTracker

type SessionTracker interface {
	ID(ctx *Context) (id string)
	SetID(ctx *Context, id string)
	DeleteID(ctx *Context)
}

///////////////////////////////////////////////////////////////////////////////
// CookieSessionTracker

var SessionIDCookie = "session"

// http://en.wikipedia.org/wiki/HTTP_cookie
type CookieSessionTracker struct {
}

func (*CookieSessionTracker) ID(ctx *Context) string {
	id, _ := ctx.Request.SiteCookie(SessionIDCookie)
	return id
}

func (*CookieSessionTracker) SetID(ctx *Context, id string) {
	ctx.Response.SetSiteCookie(SessionIDCookie, id)
}

func (*CookieSessionTracker) DeleteID(ctx *Context) {
	ctx.Response.DeleteSiteCookie(SessionIDCookie)
}
