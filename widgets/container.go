package widgets

import (
	option "github.com/fengdotdev/golibs-options"
	"github.com/fengdotdev/golibs-staticpages/interfaces"
	"github.com/fengdotdev/golibs-staticpages/models/webmodels/htmlmodels"
	"github.com/fengdotdev/golibs-staticpages/typesdef"
)

type Container struct {
	children         Widget
	containerOptions option.Option[ContainerOptions]
	classBaseOpt     typesdef.ClassBaseOpt
}

func NewContainer(children Widget, containerOptions ContainerOptions, classBaseOpt typesdef.ClassBaseOpt) *Container {
	return &Container{
		children:         children,
		containerOptions: option.NewSome(containerOptions),
		classBaseOpt:     classBaseOpt,
	}
}

func NewEmptyContainer() *Container {
	return &Container{
		children:         nil,
		containerOptions: option.NewNone[ContainerOptions](),
		classBaseOpt:     "",
	}
}

type ContainerOptions struct {
	Id string
}

// WidgetMisc
func (c *Container) HaveChildren() bool {
	return c.children != nil
}
func (c *Container) NumbOfChildren() int {
	if c.children != nil {
		return 1
	}
	return 0
}
func (c *Container) HaveOptions() bool {
	return c.containerOptions.IsSome()
}

// interfaces.WidgetConversions
func (c *Container) ToElementHTML() interfaces.ElementHTML {

	haveChildren := c.children != nil
	var attributes htmlmodels.Attributes = htmlmodels.NewAttributesEmpty()

	if c.classBaseOpt != "" {
		attributes.AppendAttribute("class", c.classBaseOpt.String())
	}

	if c.containerOptions.IsSome() {
		options := c.containerOptions.UnwrapOr(ContainerOptions{})
		if options.Id != "" {
			attributes.AddAttribute("id", options.Id)
		}
	}

	if haveChildren {
		var output htmlmodels.HTMLElement = htmlmodels.NewHTMLElement("div", attributes, []htmlmodels.HTMLElementInterface{c.children.ToElementHTML()}, "")
		return &output
	}

	var output htmlmodels.HTMLElement = htmlmodels.NewHTMLElement("div", attributes, []htmlmodels.HTMLElementInterface{}, "")
	return &output
}

// interfaces.Render
func (c *Container) RenderHTML() typesdef.HTML {
	element := c.ToElementHTML()
	return element.RenderHTML()
}

// traits.JSONTrait
func (c *Container) ToJSON() (string, error) {
	panic("implement me")
	return "", nil
}
func (c *Container) FromJSON(s string) error {
	panic("implement me")
	return nil
}
