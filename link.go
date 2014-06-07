package view

// type LinkModel interface {
// 	URLGetter
// 	LinkContent(ctx *Context) View
// 	LinkTitle(ctx *Context) string
// 	LinkRel(ctx *Context) string
// }

// Link represents an HTML <a> or <link> element depending on UseLinkTag.
// Content and title of the Model will only be rendered for <a>.
type Link struct {
	URLGetter
	ID         string
	Class      string
	Content    View
	Title      string
	Rel        string
	NewWindow  bool
	UseLinkTag bool
}

func (self *Link) Render(ctx *Context) (err error) {
	if self.UseLinkTag {
		ctx.Response.XML.OpenTag("link")
	} else {
		ctx.Response.XML.OpenTag("a")
	}
	ctx.Response.XML.AttribIfNotDefault("id", self.ID)
	ctx.Response.XML.AttribIfNotDefault("class", self.Class)
	if self.NewWindow {
		ctx.Response.XML.Attrib("target", "_blank")
	}
	ctx.Response.XML.Attrib("href", self.URL(ctx))
	ctx.Response.XML.AttribIfNotDefault("rel", self.Rel)
	if self.UseLinkTag {
		ctx.Response.XML.CloseTag() // link
	} else {
		ctx.Response.XML.AttribIfNotDefault("title", self.Title)
		if self.Content != nil {
			err = self.Content.Render(ctx)
		}
		ctx.Response.XML.CloseTagAlways() // a
	}
	return err
}
