package view

// Views implements the View interface for a slice of views.
type Views []View

func (views Views) Render(ctx *Context) (err error) {
	for _, view := range views {
		if view != nil {
			err = view.Render(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (views *Views) Insert(index int, view View) {
	*views = append(*views, nil)
	copy((*views)[index+1:], (*views)[index:])
	(*views)[index] = view
}

func (views *Views) Remove(index int) {
	copy((*views)[index:], (*views)[index+1:])
	*views = (*views)[:len(*views)-1]
}

func AsViews(values ...interface{}) Views {
	views := make(Views, len(values))
	for i, value := range values {
		views[i] = asView(value)
	}
	return views	
}