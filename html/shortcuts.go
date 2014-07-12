package html

import (
	"html"

	"github.com/gostart/view"
)

var (
	// BR creates <br/>
	BR = view.String("<br/>")

	// HR creates <hr/>
	HR = view.String("<hr/>")

	// DIVClearBoth creates <div style='clear:both'></div>
	DIVClearBoth = view.String("<div style='clear:both'></div>")
)

// Escape HTML-escapes a string.
func Escape(text string) view.String {
	return view.String(html.EscapeString(text))
}

// STYLE creates <style>css</style>
func STYLE(css string) view.String {
	return view.Printf("<style>%s</style>", css)
}

// SCRIPT creates <script>javascript</script>
func SCRIPT(javascript string) view.String {
	return view.Printf("<script>%s</script>", javascript)
}

// ScriptLink creates <script src='url'></script>
func ScriptLink(url string) view.String {
	return view.Printf("<script src='%s'></script>", url)
}
