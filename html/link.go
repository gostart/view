package html

import (
	"github.com/gostart/view"
)

// Link represents an HTML <a> or <link> element depending on UseLinkTag.
type Link struct {
	view.URL
	UseLinkTag bool
	ID         string
	Class      string
	Title      string
	Rel        string
	NewWindow  bool
	OnClick    string
	// Name       string
	Content view.View
}

func (link *Link) Render(ctx *view.Context) (err error) {
	if link.UseLinkTag {
		ctx.Response.Out("<link")
	} else {
		ctx.Response.Out("<a")
	}
	if link.ID != "" {
		WriteAttrib(ctx.Response, "id", link.ID)
	}
	if link.Class != "" {
		WriteAttrib(ctx.Response, "class", link.Class)
	}
	if link.Title != "" {
		WriteAttrib(ctx.Response, "title", link.Title)
	}
	if link.Rel != "" {
		WriteAttrib(ctx.Response, "rel", link.Rel)
	}
	if link.NewWindow {
		WriteAttrib(ctx.Response, "target", "_blank")
	}
	if link.OnClick != "" {
		WriteAttrib(ctx.Response, "onclick", link.OnClick)
	}
	if link.NewWindow {
		WriteAttrib(ctx.Response, "href", link.GetURL(ctx))
	}
	ctx.Response.Out(">")
	if link.Content != nil {
		err = link.Content.Render(ctx)
		if err != nil {
			return err
		}
	}
	if link.UseLinkTag {
		ctx.Response.Out("</link>")
	} else {
		ctx.Response.Out("</a>")
	}
	return nil
}

func (link *Link) GetID() string {
	return link.ID
}

func (link *Link) SetID(id string) {
	link.ID = id
}

// A creates <a href="url">content</a>
func A(url view.URL, content ...interface{}) *Link {
	return &Link{URL: url, Content: view.AsView(content...)}
}

// A_nofollow creates <a href="url" rel="nofollow">content</a>
func A_nofollow(url view.URL, content ...interface{}) *Link {
	return &Link{URL: url, Content: view.AsView(content...), Rel: "nofollow"}
}

// A_blank creates <a href="url" target="_blank">content</a>
func A_blank(url view.URL, content ...interface{}) *Link {
	return &Link{URL: url, Content: view.AsView(content...), NewWindow: true}
}

// A_blank_nofollow creates <a href="url" target="_blank" rel="nofollow">content</a>
func A_blank_nofollow(url view.URL, content ...interface{}) *Link {
	return &Link{URL: url, Content: view.AsView(content...), Rel: "nofollow", NewWindow: true}
}

// A_name creates a named anchor
func A_name(name string) view.String {
	return view.Printf("<a name='%s'></a>", name)
}

// StylesheetLink creates <link rel='stylesheet' href='url'>
func StylesheetLink(url view.URL) *Link {
	return &Link{
		UseLinkTag: true,
		URL:        url,
		Rel:        "stylesheet",
	}
}
