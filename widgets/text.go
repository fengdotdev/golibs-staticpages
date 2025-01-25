package widgets

type Text struct {
	text         string
	options      TextOptions
	classBaseOpt string
}

// if the environment is not using tailwindcss or bootstrap will default to TextOptions
func NewText(text string, options TextOptions, classBaseOpt string) *Text {
	return &Text{
		text:         text,
		options:      options,
		classBaseOpt: classBaseOpt,
	}
}


// NewTextWithOptions creates a new Text with options
func NewTextWithOptions(text string, options TextOptions) *Text {
	return &Text{
		text:         text,
		options:      options,
		classBaseOpt: "",
	}
}

// NewTextWithClass creates a new Text with a class base option
// u can use tailwindcss classes like "text-red-500" or "text-lg" or
// bootstrap classes like "text-danger" or "text-lg"
func NewTextWithClass(text string, classBaseOpt string) *Text {
	return &Text{
		text:         text,
		classBaseOpt: classBaseOpt,
	}
}

func (t *Text) SetClassBaseOpt(classBaseOpt string) {
	t.classBaseOpt = classBaseOpt
}

type TextOptions struct {
	Font  string
	Color string
	Size  int
}
