package html

import (
	"fmt"
	"sync"

	"github.com/gostart/view"
)

type ViewWithID interface {
	view.View
	GetID() string
	SetID(string)
}

var (
	idCounter      uint64
	idCounterMutex sync.Mutex
)

func UniqueID() string {
	idCounterMutex.Lock()
	defer idCounterMutex.Unlock()

	return fmt.Sprint("ID", idCounter)
}

// LABEL creates a Label for target and returns it together with target.
func LABEL(label interface{}, target ViewWithID) view.Views {
	id := target.GetID()
	if id == "" {
		id = UniqueID()
		target.SetID(id)
	}
	return view.Views{
		Element("label", "for", id, view.AsView(label)),
		target,
	}
}
