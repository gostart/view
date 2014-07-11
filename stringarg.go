package view

type pathArg interface {
	pathArg()
}

type StringArg struct {
}

func (StringArg) pathArg() {}
