package typesdef

type ClassBaseOpt string

func TailWindCCSClass(class string) ClassBaseOpt {
	return ClassBaseOpt(class)
}

func BootstrapClass(class string) ClassBaseOpt {
	return ClassBaseOpt(class)
}

func NewClassBaseOpt(class string) ClassBaseOpt {
	return ClassBaseOpt(class)
}

func (c ClassBaseOpt) String() string {
	return string(c)
}
