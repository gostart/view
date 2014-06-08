package view

var Site = struct {
	Endpoint
	Project struct {
		StringArg
		Endpoint
		Law  Endpoint
		Calc Endpoint
	}
}{}

func init() {
	// Site.Project.Law =
}

type Endpoint struct {
}
