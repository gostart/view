package view

import (
	"bytes"
	"container/heap"
	"fmt"
	"hash/crc32"
	"mime"
	"net/http"
	"strings"
	"time"
)

func newResponse(server *Server, responseWriter http.ResponseWriter) *Response {
	return &Response{
		server: server,
		writer: responseWriter,
	}
}

type Response struct {
	bytes.Buffer

	server *Server
	writer http.ResponseWriter

	dynamicStyle       dependencyHeap
	dynamicHeadScripts dependencyHeap
	dynamicScripts     dependencyHeap
}

// GetBytesAndReset() returns a copy of the response buffer and resets it.
func (response *Response) GetBytesAndReset() []byte {
	b := response.Buffer.Bytes()
	response.Buffer = bytes.Buffer{}
	return b
}

// GetStringAndReset() returns a copy of the response buffer and resets it.
func (response *Response) GetStringAndReset() string {
	s := response.Buffer.String()
	response.Buffer.Reset()
	return s
}

func (response *Response) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(response.writer, a...)
}

func (response *Response) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(response.writer, format, a...)
}

func (response *Response) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(response.writer, a...)
}

// Out does the same as WriteString, except that it returns
// *Response to allow call chaining.
func (response *Response) Out(s string) *Response {
	_, err := response.WriteString(s)
	if err != nil {
		panic(err)
	}
	return response
}

func (response *Response) Header() http.Header {
	return response.writer.Header()
}

func (response *Response) WriteHeader(code int) {
	response.writer.WriteHeader(code)

}

func (response *Response) SetSiteCookie(name, value string) {
	cookieCipher := response.server.CookieCipher
	if cookieCipher != nil {
		value = string(cookieCipher.Encrypt([]byte(value)))
	}
	http.SetCookie(response.writer, &http.Cookie{Name: name, Value: value, Path: "/"})
}

func (response *Response) SetSiteCookieExpires(name, value string, expires time.Time) {
	cookieCipher := response.server.CookieCipher
	if cookieCipher != nil {
		value = string(cookieCipher.Encrypt([]byte(value)))
	}
	http.SetCookie(response.writer, &http.Cookie{Name: name, Value: value, Path: "/", Expires: expires})
}

func (response *Response) SetSiteCookieBytes(name string, value []byte) {
	var valueStr string
	cookieCipher := response.server.CookieCipher
	if cookieCipher != nil {
		valueStr = cookieCipher.Encrypt(value)
	} else {
		valueStr = string(value)
	}
	http.SetCookie(response.writer, &http.Cookie{Name: name, Value: valueStr, Path: "/"})
}

func (response *Response) SetSiteCookieBytesExpires(name string, value []byte, expires time.Time) {
	var valueStr string
	cookieCipher := response.server.CookieCipher
	if cookieCipher != nil {
		valueStr = cookieCipher.Encrypt(value)
	} else {
		valueStr = string(value)
	}
	http.SetCookie(response.writer, &http.Cookie{Name: name, Value: valueStr, Path: "/", Expires: expires})
}

func (response *Response) DeleteSiteCookie(name string) {
	http.SetCookie(response.writer, &http.Cookie{Name: name, Value: "delete", Path: "/", MaxAge: -1})
}

func (response *Response) Respond(code int, body string) {
	response.WriteHeader(code)
	response.Reset() // reset buffer so it will only contain body
	response.WriteString(body)
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
