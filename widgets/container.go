package widgets

import "github.com/fengdotdev/golibs-staticpages/typesdef"

type Container struct {
	children         Widget
	containerOptions ContainerOptions
	classBaseOpt     typesdef.ClassBaseOpt
}

func NewContainer(children Widget, containerOptions ContainerOptions, classBaseOpt typesdef.ClassBaseOpt) *Container {
	return &Container{
		children:         children,
		containerOptions: containerOptions,
		classBaseOpt:     classBaseOpt,
	}
}

type ContainerOptions struct {
}
