package html

import (
	"bytes"
	"fmt"
	"html"
	"io"

	"github.com/gostart/view"
)

func Element(tag string, args ...interface{}) view.View {
	lenArgs := len(args)

	var buf bytes.Buffer
	buf.WriteByte('<')
	buf.WriteString(tag)
	for i := 0; i < lenArgs; i += 2 {
		writeAttrib(&buf, args[i], args[i+1])
	}
	buf.WriteByte('>')

	if lenArgs&1 == 0 {
		// If there is no uneven last argument as view,
		// close tag and return buf bytes.

		buf.WriteString("</")
		buf.WriteString(tag)
		buf.WriteByte('>')

		return view.String(buf.String())
	}

	return &elem{
		open:  buf.String(),
		view:  view.AsView(args[lenArgs-1]),
		close: "</" + tag + ">",
	}
}

func writeAttrib(writer io.Writer, name, value interface{}) {
	fmt.Fprint(writer, " ", name, "='", html.EscapeString(fmt.Sprint(value)), "'")
}

type elem struct {
	open  string
	view  view.View
	close string
}

func (e *elem) Render(ctx *view.Context) (err error) {
	ctx.Response.Out(e.open)
	if e.view != nil {
		err = e.view.Render(ctx)
	}
	if err == nil {
		ctx.Response.Out(e.close)
	}
	return err
}

///////////////////////////////////////////////////////////////////////////////
// Functions for elements

// H1 creates <h1>content</h1>
func H1(content ...interface{}) view.View {
	return Element("h1", view.AsView(content...))
}

// H2 creates <h2>content</h2>
func H2(content ...interface{}) view.View {
	return Element("h2", view.AsView(content...))
}

// H3 creates <h3>content</h3>
func H3(content ...interface{}) view.View {
	return Element("h3", view.AsView(content...))
}

// H4 creates <h4>content</h4>
func H4(content ...interface{}) view.View {
	return Element("h4", view.AsView(content...))
}

// H5 creates <h5>content</h5>
func H5(content ...interface{}) view.View {
	return Element("h5", view.AsView(content...))
}

// H6 creates <h6>content</h6>
func H6(content ...interface{}) view.View {
	return Element("h6", view.AsView(content...))
}

// P creates <p>content</p>
func P(content ...interface{}) view.View {
	return Element("p", view.AsView(content...))
}

// B creates <b>content</b>
func B(content ...interface{}) view.View {
	return Element("b", view.AsView(content...))
}

// I creates <i>content</i>
func I(content ...interface{}) view.View {
	return Element("i", view.AsView(content...))
}

// Q creates <q>content</q>
func Q(content ...interface{}) view.View {
	return Element("q", view.AsView(content...))
}

// DEL creates <del>content</del>
func DEL(content ...interface{}) view.View {
	return Element("del", view.AsView(content...))
}

// EM creates <em>content</em>
func EM(content ...interface{}) view.View {
	return Element("em", view.AsView(content...))
}

// STRONG creates <strong>content</strong>
func STRONG(content ...interface{}) view.View {
	return Element("strong", view.AsView(content...))
}

// DFN creates <dfn>content</dfn>
func DFN(content ...interface{}) view.View {
	return Element("dfn", view.AsView(content...))
}

// CODE creates <code>content</code>
func CODE(content ...interface{}) view.View {
	return Element("code", view.AsView(content...))
}

// PRE creates <pre>content</pre>
func PRE(content ...interface{}) view.View {
	return Element("pre", view.AsView(content...))
}

// SECTION creates <sections class="class">content</section>
func SECTION(class string, content ...interface{}) view.View {
	return Element("section", "class", class, view.AsView(content...))
}

// ABBR creates <abbr title="longTitle">abbreviation</abbr>
func ABBR(longTitle, abbreviation string) view.View {
	return Element("abbr", "title", longTitle, Escape(abbreviation))
}

// RSSLink creates <link rel='alternate' type='application/rss+xml' title='title' href='url'>
func RSSLink(title string, url view.URL) view.View {
	return &elem{
		open:  fmt.Sprintf("<link rel='alternate' type='application/rss+xml' title='%s' href='", title),
		view:  view.URLView{url},
		close: "'>",
	}
}
