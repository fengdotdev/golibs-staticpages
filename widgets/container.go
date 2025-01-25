package widgets

type Container struct {
	children Widget
}



func NewContainer(children Widget) *Container {
	return &Container{
		children: children,
	}
}
