package view

import (
	"bytes"
	"container/heap"
	"fmt"
	"hash/crc32"
	"mime"
	"net/http"
	"strings"
)

func newResponse(responseWriter http.ResponseWriter) *Response {
	response := &Response{
		ResponseWriter: responseWriter,
	}
	response.PushBody()
	return response
}

type responseBody struct {
	buffer *bytes.Buffer
	xml    *XMLWriter
}

type Response struct {
	http.ResponseWriter

	// Session *Session

	bodyStack []responseBody

	XML *XMLWriter // XML allowes the Response to be used as XMLWriter

	dynamicStyle       dependencyHeap
	dynamicHeadScripts dependencyHeap
	dynamicScripts     dependencyHeap
}

// PushBody pushes the buffer of the response body on a stack
// and sets a new empty buffer.
// This can be used to render intermediate text results.
// Note: Only the response body is pushed, all other state changes
// like setting headers will affect the final response.
func (response *Response) PushBody() {
	var b responseBody
	b.buffer = new(bytes.Buffer)
	b.xml = NewXMLWriter(b.buffer)
	response.bodyStack = append(response.bodyStack, b)
	response.XML = b.xml
}

// PopBody pops the buffer of the response body from the stack
// and returns its content.
func (response *Response) PopBody() (bufferData []byte) {
	last := len(response.bodyStack) - 1
	bufferData = response.bodyStack[last].buffer.Bytes()
	response.bodyStack = response.bodyStack[0:last]
	response.XML = response.bodyStack[last-1].xml
	return bufferData
}

// PopBodyString pops the buffer of the response body from the stack
// and returns its content as string.
func (response *Response) PopBodyString() (bufferData string) {
	return string(response.PopBody())
}

func (response *Response) Write(p []byte) (n int, err error) {
	return response.XML.Write(p)
}

func (response *Response) WriteByte(c byte) error {
	_, err := response.XML.Write([]byte{c})
	return err
}

func (response *Response) Print(s string) *Response {
	_, err := response.XML.Write([]byte(s))
	if err != nil {
		panic(err)
	}
	return response
}

func (response *Response) Printf(format string, args ...interface{}) (n int, err error) {
	return fmt.Fprintf(response.XML, format, args...)
}

func (response *Response) String() string {
	return response.bodyStack[len(response.bodyStack)-1].buffer.String()
}

func (response *Response) Bytes() []byte {
	return response.bodyStack[len(response.bodyStack)-1].buffer.Bytes()
}

func (response *Response) SetSecureCookie(name string, val string, age int64, path string) {
	panic("not implemented")
	/// todo: ";HttpOnly"
}

func (response *Response) Respond(code int, body string) {
	response.WriteHeader(code)
	response.Print(body)
}

func (response *Response) RespondPlainText(code int, body string) {
	response.SetContentTypePlainText()
	response.Respond(code, body)
}

func (response *Response) SetContentType(contentType string) {
	response.Header().Set("Content-Type", contentType)
}

func (response *Response) SetContentTypeByExt(ext string) {
	response.SetContentType(mime.TypeByExtension(ext))
}

func (response *Response) SetContentTypePlainText() {
	response.SetContentType("text/plain; charset=utf-8")
}

func (response *Response) SetContentTypeHTML() {
	response.SetContentType("text/html; charset=utf-8")
}

func (response *Response) SetContentTypeXML() {
	response.SetContentType("application/xml; charset=utf-8")
}

func (response *Response) SetContentTypeJSON() {
	response.SetContentType("application/json; charset=utf-8")
}

// SetContentTypeAttachment makes the webbrowser open a
// "Save As.." dialog for the response.
func (response *Response) SetContentTypeAttachment(filename string) {
	response.Header().Set("Content-Type", "application/x-unknown")
	response.Header().Set("Content-Disposition", "attachment;filename="+filename)
}

// RequireStyle adds dynamic CSS content to the page.
// Multiple dynamic entries will be sorted by priority.
// Dynamic CSS will be inserted after the regular CSS of the page.
// If css does not start with "<style",
// then the css string will be wrapped with a style tag.
//
// Use this feature to dynamically add CSS to the page if the
// HTML content requires it.
func (response *Response) RequireStyle(css string, priority int) {
	if response.dynamicStyle == nil {
		response.dynamicStyle = newDependencyHeap()
	}
	if strings.Index(strings.ToLower(css), "<style") != 0 {
		css = "<style>" + css + "</style>"
	}
	response.dynamicStyle.AddIfNew(css, priority)
}

// RequireStyleURL adds a dynamic CSS link to the page.
// Multiple dynamic entries will be sorted by priority.
// Dynamic CSS will be inserted after the regular CSS of the page.
//
// Use this feature to dynamically add CSS to the page if the
// HTML content requires it.
func (response *Response) RequireStyleURL(url string, priority int) {
	if response.dynamicStyle == nil {
		response.dynamicStyle = newDependencyHeap()
	}
	response.dynamicStyle.AddIfNew("<link rel='stylesheet' href='"+url+"'>", priority)
}

// RequireHeadScript adds dynamic JavaScript to the page.
// Multiple dynamic entries will be sorted by priority.
// The dynamic JavaScript will be inserted after the regular
// head-scripts of the page.
// If script does not start with "<script",
// then the script string will be wrapped with a script tag.
//
// Use this feature to dynamically add JavaScript to
// the page if the HTML content requires it.
func (response *Response) RequireHeadScript(script string, priority int) {
	if response.dynamicHeadScripts == nil {
		response.dynamicHeadScripts = newDependencyHeap()
	}
	if strings.Index(strings.ToLower(script), "<script") != 0 {
		script = "<script>" + script + "</script>"
	}
	response.dynamicHeadScripts.AddIfNew(script, priority)
}

// RequireHeadScriptURL adds a dynamic JavaScript link to the page.
// Multiple dynamic entries will be sorted by priority.
// The dynamic JavaScript will be inserted after the regular
// head-scripts of the page.
//
// Use this feature to dynamically add JavaScript to
// the page if the HTML content requires it.
func (response *Response) RequireHeadScriptURL(url string, priority int) {
	if response.dynamicHeadScripts == nil {
		response.dynamicHeadScripts = newDependencyHeap()
	}
	response.dynamicHeadScripts.AddIfNew("<script src='"+url+"'></script>", priority)
}

// RequireScript adds dynamic JavaScript to the page.
// Multiple dynamic entries will be sorted by priority.
// The dynamic JavaScript will be inserted after the regular
// scripts near the end of the page.
// If script does not start with "<script",
// then the script string will be wrapped with a script tag.
//
// Use this feature to dynamically add JavaScript to
// the page if the HTML content requires it.
func (response *Response) RequireScript(script string, priority int) {
	if response.dynamicScripts == nil {
		response.dynamicScripts = newDependencyHeap()
	}
	if strings.Index(strings.ToLower(script), "<script") != 0 {
		script = "<script>" + script + "</script>"
	}
	response.dynamicScripts.AddIfNew(script, priority)
}

// RequireScriptURL adds a dynamic JavaScript link to the page.
// Multiple dynamic entries will be sorted by priority.
// The dynamic JavaScript will be inserted after the regular
// scripts near the end of the page.
//
// Use this feature to dynamically add JavaScript to
// the page if the HTML content requires it.
func (response *Response) RequireScriptURL(url string, priority int) {
	if response.dynamicScripts == nil {
		response.dynamicScripts = newDependencyHeap()
	}
	response.dynamicScripts.AddIfNew("<script src='"+url+"'></script>", priority)
}

///////////////////////////////////////////////////////////////////////////////
// dependencyHeap

func newDependencyHeap() dependencyHeap {
	dh := make(dependencyHeap, 0, 1)
	heap.Init(&dh)
	return dh
}

type dependencyHeapItem struct {
	text     string
	hash     uint32
	priority int
}

type dependencyHeap []dependencyHeapItem

func (response *dependencyHeap) Len() int {
	return len(*response)
}

func (response *dependencyHeap) Less(i, j int) bool {
	return (*response)[i].priority < (*response)[j].priority
}

func (response *dependencyHeap) Swap(i, j int) {
	(*response)[i], (*response)[j] = (*response)[j], (*response)[i]
}

func (response *dependencyHeap) Push(item interface{}) {
	*response = append(*response, item.(dependencyHeapItem))
}

func (response *dependencyHeap) Pop() interface{} {
	end := len(*response) - 1
	item := (*response)[end]
	*response = (*response)[:end]
	return item
}

func (response *dependencyHeap) AddIfNew(text string, priority int) {
	hash := crc32.ChecksumIEEE([]byte(text))
	for i := range *response {
		if (*response)[i].hash == hash {
			return // text is not new
		}
	}
	heap.Push(response, dependencyHeapItem{text, hash, priority})
}

func (response *dependencyHeap) String() string {
	if response == nil {
		return ""
	}
	var buf bytes.Buffer
	for i := range *response {
		buf.WriteString((*response)[i].text)
	}
	return buf.String()
}
