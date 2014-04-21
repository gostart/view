package view

// Views implements the View interface for a slice of views.
type Views []View

func (self Views) Render(ctx *Context) (err error) {
	for _, view := range self {
		if view != nil {
			err = view.Render(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
