package widgets

import (
	"errors"

	option "github.com/fengdotdev/golibs-options"
	"github.com/fengdotdev/golibs-staticpages/typesdef"
)

type Text struct {
	text         string
	options      option.Option[TextOptions]
	classBaseOpt string
}

func NewSimpleText(text string) *Text {
	return &Text{
		text:         text,
		options:      option.NewNone[TextOptions](),
		classBaseOpt: "",
	}
}

// if the environment is not using tailwindcss or bootstrap will default to TextOptions
func NewText(text string, options TextOptions, classBaseOpt string) *Text {
	return &Text{
		text:         text,
		options:      option.NewSome(options),
		classBaseOpt: classBaseOpt,
	}
}

// NewTextWithOptions creates a new Text with options
func NewTextWithOptions(text string, options TextOptions) *Text {
	return &Text{
		text:         text,
		options:      option.NewSome(options),
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
		options:      option.NewNone[TextOptions](),
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

func (t *Text) GetText() string {
	return t.text
}

// return TextOptions if set, else error
func (t *Text) GetOptions() (TextOptions, error) {
	if t.options.IsNone() {
		return TextOptions{}, errors.New("options is not set")
	}
	return t.options.UnwrapOr(TextOptions{}), nil

}

func (t *Text) RenderHTML() typesdef.HTML {
	if t.classBaseOpt != "" && t.options.IsNone() {
		s := "<p class='" + t.classBaseOpt + "'>" + t.text + "</p>"
		return *typesdef.NewHTML(s)
	}
	s := "<p>" + t.text + "</p>"
	return *typesdef.NewHTML(s)
}

func (t *Text) ToJSON() (string, error) {
	panic("Not implemented")
	return `{"text": "` + t.text + `"}`, nil
}

func (t *Text) FromJSON(s string) error {
	panic("Not implemented")
	t.text = s
	return nil
}
