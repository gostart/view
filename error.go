package view

func NewError(err error) *Error {
	return &Error{err}
}

// Error wraps an error and returns it from the Render() method
// instead of rendering anything.
type Error struct {
	error
}

func (self *Error) Render(ctx *Context) (err error) {
	return self.error
}
