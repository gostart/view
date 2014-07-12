package view

type Pointer struct {
	Ptr *View
}

func (pointer Pointer) Render(ctx *Context) error {
	if pointer.Ptr == nil || *pointer.Ptr == nil {
		return nil
	}
	return (*pointer.Ptr).Render(ctx)
}
