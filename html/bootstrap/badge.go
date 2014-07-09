package bootstrap

import (
	"github.com/gostart/view"
)

func NewBadge(class string, content view.View) *Label {
	return &Label{Class: LabelClass(class), Content: content}
}
