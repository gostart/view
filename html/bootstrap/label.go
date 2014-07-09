package bootstrap

import (
	"github.com/gostart/view"
)

type LabelClass string

func (class LabelClass) New(content view.View) *Label {
	return &Label{Class: class, Content: content}
}

const (
	LabelDefault LabelClass = "label-default"
	LabelPrimary LabelClass = "label-primary"
	LabelSuccess LabelClass = "label-success"
	LabelInfo    LabelClass = "label-info"
	LabelWarning LabelClass = "label-warning"
	LabelDanger  LabelClass = "label-danger"
)

type Label struct {
	Class   LabelClass
	Content view.View
}

func (label *Label) Render(ctx *view.Context) (err error) {
	ctx.Response.Print("<span class='label ").Print(string(label.Class)).Print("'>")
	if label.Content != nil {
		err = label.Content.Render(ctx)
	}
	ctx.Response.Print("</span>")
	return err
}
