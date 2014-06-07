package view

type List struct {
	ID          string
	Class       string
	Items       []View
	Ordered     bool
	OrderOffset uint
}

func (self *List) Render(ctx *Context) (err error) {
	if self.Ordered {
		ctx.Response.XML.OpenTag("ol")
		ctx.Response.XML.Attrib("start", self.OrderOffset+1)
	} else {
		ctx.Response.XML.OpenTag("ul")
	}
	ctx.Response.XML.AttribIfNotDefault("id", self.ID)
	ctx.Response.XML.AttribIfNotDefault("class", self.Class)

	for i, view := range self.Items {
		ctx.Response.XML.OpenTag("li")
		if self.ID != "" {
			ctx.Response.XML.Attrib("id", self.ID, "_", i)
		}
		if view != nil {
			err = view.Render(ctx)
			if err != nil {
				return err
			}
		}
		ctx.Response.XML.CloseTagAlways() // li
	}

	ctx.Response.XML.CloseTagAlways() // ol/ul
	return nil
}
