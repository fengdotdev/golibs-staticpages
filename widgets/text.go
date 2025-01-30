package widgets

import (
	"errors"

	option "github.com/fengdotdev/golibs-options"
	"github.com/fengdotdev/golibs-staticpages/interfaces"
	"github.com/fengdotdev/golibs-staticpages/models/webmodels/htmlmodels"
	"github.com/fengdotdev/golibs-staticpages/typesdef"
)

type Text struct {
	text         string
	options      option.Option[TextOptions]
	classBaseOpt typesdef.ClassBaseOpt
}

func NewSimpleText(text string) *Text {
	return &Text{
		text:         text,
		options:      option.NewNone[TextOptions](),
		classBaseOpt: "",
	}
}

// if the environment is not using tailwindcss or bootstrap will default to TextOptions
func NewText(text string, options TextOptions, classBaseOpt typesdef.ClassBaseOpt) *Text {
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
		classBaseOpt: typesdef.NewClassBaseOpt(classBaseOpt),
		options:      option.NewNone[TextOptions](),
	}
}

func (t *Text) SetClassBaseOpt(classBaseOpt string) {
	t.classBaseOpt = typesdef.NewClassBaseOpt(classBaseOpt)
}

type TextOptions struct {
	Id    string
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

//WidgetMisc

func (t *Text) NumbOfChildren() int {
	return 0
}

func (t *Text) HaveChildren() bool {
	return false
}

func (t *Text) HaveOptions() bool {
	return t.options.IsSome()
}

// WidgetConversions
func (t *Text) ToElementHTML() interfaces.ElementHTML {

	haveChildren := false

	if haveChildren {
		panic("Not implemented")
	}

	var attributes htmlmodels.Attributes = htmlmodels.NewAttributesEmpty()

	if t.options.IsSome() {
		options := t.options.UnwrapOr(TextOptions{})
		if options.Id != "" {
			attributes.AddAttribute("id", options.Id)
		}

	}
	if t.classBaseOpt != "" {
		attributes.AppendAttribute("class", t.classBaseOpt.String())
	}
	var output htmlmodels.HTMLElement = htmlmodels.NewHTMLElement("p", attributes, []htmlmodels.HTMLElementInterface{}, t.text)

	return &output
}

// interfaces.Render
func (t *Text) RenderHTML() typesdef.HTML {
	element := t.ToElementHTML()
	return element.RenderHTML()
}

// traits.JSONTrait
func (t *Text) ToJSON() (string, error) {
	panic("Not implemented")
	return `{"text": "` + t.text + `"}`, nil
}

func (t *Text) FromJSON(s string) error {
	panic("Not implemented")
	t.text = s
	return nil
}
