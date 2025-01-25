package widgets

type Button struct {
	text         string
	action       string
	options      ButtonOptions
	classBaseOpt string
}

type ButtonOptions struct {
	Font  string
	Color string
	Size  int
}
