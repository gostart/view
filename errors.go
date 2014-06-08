package view

import (
	"fmt"
)

type ViewError interface {
	View
	error
}

var (
	AuthorizationRequired401 ViewError = &plainTextResponse{401, "401 authorization required"}
	Forbidden403             ViewError = &plainTextResponse{403, "403 forbidden"}
	NotFound404              ViewError = &plainTextResponse{404, "404 page not found"}
	NotModified304           ViewError = &response{304, "not modfied"}
)

func InternalServerError500(err error) ViewError {
	return &internalServerError500{err}
}

func RedirectPermanently301(url URLGetter) ViewError {
	return &redirect{301, url}
}

func RedirectTemporary302(url URLGetter) ViewError {
	return &redirect{302, url}
}

type internalServerError500 struct {
	error
}

func (self *internalServerError500) Render(ctx *Context) error {
	message := "500 internal server error"
	if self.error != nil && Config.Debug.Mode {
		message += "\n\n" + self.error.Error()
	}
	ctx.Response.RespondPlainText(500, message)
	return nil
}

type redirect struct {
	code int
	url  URLGetter
}

func (self *redirect) Render(ctx *Context) error {
	ctx.Response.Respond(self.code, self.url.URL(ctx))
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
