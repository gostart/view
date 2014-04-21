package view

type HTML string

func (self HTML) Render(ctx *Context) (err error) {
	_, err = ctx.Response.Write([]byte(self))
	return err
}

// A_name creates a named anchor
func A_name(name string) HTML {
	return Printf("<a name='%s'></a>", name)
}

// STYLE creates <style>css</style>
func STYLE(css string) HTML {
	return Printf("<style>%s</style>", css)
}

// SCRIPT creates <script>javascript</script>
func SCRIPT(javascript string) HTML {
	return Printf("<script>%s</script>", javascript)
}

// ScriptLink creates <script src='url'></script>
func ScriptLink(url string) HTML {
	return Printf("<script src='%s'></script>", url)
}

// DivClearBoth creates <div style='clear:both'></div>
func DivClearBoth() HTML {
	return HTML("<div style='clear:both'></div>")
}

// BR creates <br/>
func BR() HTML {
	return HTML("<br/>")
}

// HR creates <hr/>
func HR() HTML {
	return HTML("<hr/>")
}
