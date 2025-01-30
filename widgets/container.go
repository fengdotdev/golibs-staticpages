package widgets

import (
	option "github.com/fengdotdev/golibs-options"
	"github.com/fengdotdev/golibs-staticpages/helperfuncs"
	"github.com/fengdotdev/golibs-staticpages/interfaces"
	"github.com/fengdotdev/golibs-staticpages/models/webmodels/htmlmodels"
	"github.com/fengdotdev/golibs-staticpages/style"
	"github.com/fengdotdev/golibs-staticpages/typesdef"
)

func ContainerIdentifier(id string) string {
	name := id
	toHash := "container"
	return helperfuncs.GenerateIdDeterministicWithHASH(name, toHash)
}

type Container struct {
	identifier       string
	children         Widget
	containerOptions option.Option[ContainerOptions]
	classBaseOpt     typesdef.ClassBaseOpt
}

func containerIdentifier(id string) string {
	var identifier string
	haveId := id != ""
	if haveId {
		identifier = ContainerIdentifier(id)
		return identifier
	}
	identifier = ContainerIdentifier("container")
	return identifier
}

func NewContainer(children Widget, containerOptions ContainerOptions, classBaseOpt typesdef.ClassBaseOpt) *Container {

	return &Container{
		identifier:       containerIdentifier(containerOptions.Id),
		children:         children,
		containerOptions: option.NewSome(containerOptions),
		classBaseOpt:     classBaseOpt,
	}
}

func NewEmptyContainer() *Container {
	return &Container{
		identifier:       containerIdentifier(""),
		children:         nil,
		containerOptions: option.NewNone[ContainerOptions](),
		classBaseOpt:     "",
	}
}

type ContainerOptions struct {
	Id              string
	Size            style.Size
	Padding         style.Padding
	Margin          style.Margin
	BackgroundColor style.Color
}

// WidgetMisc
func (c *Container) Identifier() string {
	return c.identifier
}

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

//webre
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
