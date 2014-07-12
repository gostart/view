package view

import (
	"fmt"
)

type ViewError interface {
	View
	error
}

// Error returns Err when Render is called.
// Error does not implement ViewError to prevent endless loops.
type Error struct {
	Err error
}

func (err Error) Render(ctx *Context) error {
	return err.Err
}

var (
	ErrNotModified304           ViewError = &response{304, "not modfied"}
	ErrAuthorizationRequired401 ViewError = &plainTextResponse{401, "401 authorization required"}
	ErrForbidden403             ViewError = &plainTextResponse{403, "403 forbidden"}
	ErrNotFound404              ViewError = &plainTextResponse{404, "404 page not found"}
)

func ErrRedirectPermanently301(url URL) ViewError {
	return &redirect{301, url}
}

func ErrRedirectTemporary302(url URL) ViewError {
	return &redirect{302, url}
}

func ErrInternalServerError500(err error) ViewError {
	return &internalServerError500{err}
}

type internalServerError500 struct {
	error
}

func (self *internalServerError500) Render(ctx *Context) error {
	message := "500 internal server error"
	if self.error != nil && ctx.Server.Debug.Mode {
		message += "\n\n" + self.error.Error()
	}
	ctx.Response.RespondPlainText(500, message)
	return nil
}

type redirect struct {
	code int
	url  URL
}

func (self *redirect) Render(ctx *Context) error {
	ctx.Response.Respond(self.code, self.url.GetURL(ctx))
	return nil
}

func (self *redirect) Error() string {
	return fmt.Sprintf("%d redirect %#v", self.code, self.url)
}

type response struct {
	code int
	body string
}

func (self response) Render(ctx *Context) error {
	ctx.Response.Respond(self.code, self.body)
	return nil
}

func (self response) Error() string {
	return self.body
}

type plainTextResponse struct {
	code int
	body string
}

func (self *plainTextResponse) Render(ctx *Context) error {
	ctx.Response.RespondPlainText(self.code, self.body)
	return nil
}

func (self *plainTextResponse) Error() string {
	return self.body
}
