package view

import (
	"fmt"
	"html"
	"reflect"
)

// ViewOrError returns view if err is nil, or else an Error view for err.
func ViewOrError(view View, err error) View {
	if err != nil {
		return ErrInternalServerError500(err)
	} else {
		return view
	}
}

// Escape HTML escapes a string.
func Escape(text string) HTML {
	return HTML(html.EscapeString(text))
}

// Printf creates an unescaped HTML string.
func Printf(text string, args ...interface{}) HTML {
	return HTML(fmt.Sprintf(text, args...))
}

// PrintfEscape creates an escaped HTML string.
func PrintfEscape(text string, args ...interface{}) HTML {
	return Escape(fmt.Sprintf(text, args...))
}

// // LABEL creates a Label for target and returns it together with target.
// func LABEL(label interface{}, target View) Views {
// 	if target.ID() == "" {
// 		panic("Label target must have an ID")
// 	}
// 	return Views{
// 		&Label{
// 			Content: NewView(label),
// 			For:     target,
// 		},
// 		target,
// 	}
// }

// A creates <a href="url">content</a>
func A(url URLGetter, content ...interface{}) *Link {
	return &Link{URLGetter: url, Content: WrapContents(content...)}
}

// A_nofollow creates <a href="url" rel="nofollow">content</a>
func A_nofollow(url URLGetter, content ...interface{}) *Link {
	return &Link{URLGetter: url, Content: WrapContents(content...), Rel: "nofollow"}
}

// A_blank creates <a href="url" target="_blank">content</a>
func A_blank(url URLGetter, content ...interface{}) *Link {
	return &Link{URLGetter: url, Content: WrapContents(content...), NewWindow: true}
}

// A_blank_nofollow creates <a href="url" target="_blank" rel="nofollow">content</a>
func A_blank_nofollow(url URLGetter, content ...interface{}) *Link {
	return &Link{URLGetter: url, Content: WrapContents(content...), Rel: "nofollow", NewWindow: true}
}

// StylesheetLink creates <link rel='stylesheet' href='url'>
func StylesheetLink(url URLGetter) *Link {
	return &Link{
		URLGetter:  url,
		Rel:        "stylesheet",
		UseLinkTag: true,
	}
}

// RSSLink creates <link rel='alternate' type='application/rss+xml' title='title' href='url'>
func RSSLink(title string, url URLGetter) View {
	return ViewFunc(
		func(ctx *Context) error {
			href := url.URL(ctx)
			ctx.Response.Printf("<link rel='alternate' type='application/rss+xml' title='%s' href='%s'>", title, href)
			return nil
		},
	)
}

// IMG creates a HTML img element for an URL with optional width and height.
// The first int of dimensions is width, the second one height.
func IMG(url string, dimensions ...int) *Image {
	var width int
	var height int
	dimCount := len(dimensions)
	if dimCount >= 1 {
		width = dimensions[0]
		if dimCount >= 2 {
			height = dimensions[1]
		}
	}
	return &Image{Src: url, Width: width, Height: height}
}

// SECTION creates <sections class="class">content</section>
func SECTION(class string, content ...interface{}) View {
	return &Tag{Tag: "section", Class: class, Content: WrapContents(content...)}
}

// DIV creates <span class="class">content</span>
func SPAN(class string, content ...interface{}) *Span {
	return &Span{Class: class, Content: WrapContents(content...)}
}

// P creates <p>content</p>
func P(content ...interface{}) View {
	return &Tag{Tag: "p", Content: WrapContents(content...)}
}

// H1 creates <h1>content</h1>
func H1(content ...interface{}) View {
	return &Tag{Tag: "h1", Content: WrapContents(content...)}
}

// H2 creates <h2>content</h2>
func H2(content ...interface{}) View {
	return &Tag{Tag: "h2", Content: WrapContents(content...)}
}

// H3 creates <h3>content</h3>
func H3(content ...interface{}) View {
	return &Tag{Tag: "h3", Content: WrapContents(content...)}
}

// H4 creates <h4>content</h4>
func H4(content ...interface{}) View {
	return &Tag{Tag: "h4", Content: WrapContents(content...)}
}

// H5 creates <h5>content</h5>
func H5(content ...interface{}) View {
	return &Tag{Tag: "h5", Content: WrapContents(content...)}
}

// H creates <h6>content</h6>
func H6(content ...interface{}) View {
	return &Tag{Tag: "h6", Content: WrapContents(content...)}
}

// B creates <b>content</b>
func B(content ...interface{}) View {
	return &Tag{Tag: "b", Content: WrapContents(content...)}
}

// I creates <i>content</i>
func I(content ...interface{}) View {
	return &Tag{Tag: "i", Content: WrapContents(content...)}
}

// Q creates <q>content</q>
func Q(content ...interface{}) View {
	return &Tag{Tag: "q", Content: WrapContents(content...)}
}

// DEL creates <del>content</del>
func DEL(content ...interface{}) View {
	return &Tag{Tag: "del", Content: WrapContents(content...)}
}

// EM creates <em>content</em>
func EM(content ...interface{}) View {
	return &Tag{Tag: "em", Content: WrapContents(content...)}
}

// STRONG creates <strong>content</strong>
func STRONG(content ...interface{}) View {
	return &Tag{Tag: "strong", Content: WrapContents(content...)}
}

// DFN creates <dfn>content</dfn>
func DFN(content ...interface{}) View {
	return &Tag{Tag: "dfn", Content: WrapContents(content...)}
}

// CODE creates <code>content</code>
func CODE(content ...interface{}) View {
	return &Tag{Tag: "code", Content: WrapContents(content...)}
}

// PRE creates <pre>content</pre>
func PRE(content ...interface{}) View {
	return &Tag{Tag: "pre", Content: WrapContents(content...)}
}

// ABBR creates <abbr title="longTitle">abbreviation</abbr>
func ABBR(longTitle, abbreviation string) View {
	return &Tag{Tag: "abbr", Attribs: map[string]string{"title": longTitle}, Content: Escape(abbreviation)}
}

// UL is a shortcut to create an unordered list by wrapping items as HTML views.
// NewView will be called for every passed item.
//
// Example:
//   UL("red", "green", "blue")
//   UL(A(url1, "First Link"), A(url2, "Second Link"))
//
func UL(items ...interface{}) *List {
	return &List{Items: NewViews(items...)}
}

// OL is a shortcut to create an ordered list by wrapping items as HTML views.
// NewView will be called for every passed item.
//
// Example:
//   OL("red", "green", "blue")
//   OL(A(url1, "First Link"), A(url2, "Second Link"))
//
func OL(items ...interface{}) *List {
	list := UL(items...)
	list.Ordered = true
	return list
}

// NewView encapsulates content as View.
// Strings or fmt.Stringer implementations will be HTML escaped.
// View implementations will be passed through.
func NewView(content interface{}) View {
	if content == nil {
		return nil
	}
	if view, ok := content.(View); ok {
		return view
	}
	if stringer, ok := content.(fmt.Stringer); ok {
		return Escape(stringer.String())
	}
	v := reflect.ValueOf(content)
	if v.Kind() != reflect.String {
		panic(fmt.Errorf("Invalid content type: %T (must be gostart/view.View, fmt.Stringer or a string)", content))
	}
	return Escape(v.String())
}

// NewViews encapsulates multiple content arguments as views by calling NewView()
// for every one of them.
func NewViews(contents ...interface{}) Views {
	count := len(contents)
	if count == 0 {
		return nil
	}
	views := make(Views, count)
	for i, content := range contents {
		views[i] = NewView(content)
	}
	return views
}

// WrapContents encapsulates multiple content arguments as View by calling NewView()
// for every one of them.
// It is more efficient for one view because the view is passed through instead of wrapped
// with a Views slice like NewViews does.
func WrapContents(contents ...interface{}) View {
	count := len(contents)
	switch count {
	case 0:
		return nil
	case 1:
		return NewView(contents[0])
	}
	views := make(Views, count)
	for i, content := range contents {
		views[i] = NewView(content)
	}
	return views
}
