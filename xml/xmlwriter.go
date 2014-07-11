package xml

import (
	"encoding/xml"
	"fmt"
	"html"
	"io"

	"github.com/gostart/errs"
	"github.com/ungerik/go-dry"
)

func NewWriter(writer io.Writer) *Writer {
	if xmlWriter, ok := writer.(*Writer); ok {
		return xmlWriter
	}
	return &Writer{writer: writer}
}

type Writer struct {
	writer    io.Writer
	tagStack  []string
	inOpenTag bool
}

func (self *Writer) WriteXMLHeader() *Writer {
	return self.Content(xml.Header)
}

func (self *Writer) OpenTag(tag string) *Writer {
	self.finishOpenTag()

	self.writer.Write([]byte{'<'})
	self.writer.Write([]byte(tag))

	self.tagStack = append(self.tagStack, tag)
	self.inOpenTag = true

	return self
}

// value will be HTML escaped and concaternated
func (self *Writer) Attrib(name string, value ...interface{}) *Writer {
	errs.Assert(self.inOpenTag, "utils.Writer.Attrib() must be called inside of open tag")

	fmt.Fprintf(self.writer, " %s='", name)
	for _, valuePart := range value {
		str := html.EscapeString(fmt.Sprint(valuePart))
		self.writer.Write([]byte(str))
	}
	self.writer.Write([]byte{'\''})

	return self
}

func (self *Writer) AttribIfNotDefault(name string, value interface{}) *Writer {
	if dry.IsZero(value) {
		return self
	}
	return self.Attrib(name, value)
}

// AttribFlag writes a name="name" attribute if flag is true,
// else nothing will be written.
func (self *Writer) AttribFlag(name string, flag bool) *Writer {
	if flag {
		self.Attrib(name, name)
	}
	return self
}

func (self *Writer) Content(s string) *Writer {
	self.Write([]byte(s))
	return self
}

func (self *Writer) EscapeContent(s string) *Writer {
	self.Write([]byte(html.EscapeString(s)))
	return self
}

func (self *Writer) Printf(format string, args ...interface{}) *Writer {
	fmt.Fprintf(self, format, args...)
	return self
}

func (self *Writer) PrintfEscape(format string, args ...interface{}) *Writer {
	return self.EscapeContent(fmt.Sprintf(format, args...))
}

// implements io.Writer
func (self *Writer) Write(p []byte) (n int, err error) {
	self.finishOpenTag()
	return self.writer.Write(p)
}

func (self *Writer) CloseTag() *Writer {
	// this kind of sucks
	// if we can haz append() why not pop()?
	top := len(self.tagStack) - 1
	tag := self.tagStack[top]
	self.tagStack = self.tagStack[:top]

	if self.inOpenTag {
		self.writer.Write([]byte("/>"))
		self.inOpenTag = false
	} else {
		self.writer.Write([]byte("</"))
		self.writer.Write([]byte(tag))
		self.writer.Write([]byte{'>'})
	}

	return self
}

// Creates an explicit close tag, even if there is no content
func (self *Writer) CloseTagAlways() *Writer {
	self.finishOpenTag()
	return self.CloseTag()
}

func (self *Writer) finishOpenTag() {
	if self.inOpenTag {
		self.writer.Write([]byte{'>'})
		self.inOpenTag = false
	}
}

func (self *Writer) Reset() {
	if self.tagStack != nil {
		self.tagStack = self.tagStack[0:0]
	}
	self.inOpenTag = false
}
