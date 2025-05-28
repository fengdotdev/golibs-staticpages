package htmlmodels

import "github.com/fengdotdev/golibs-staticpages/typesdef"

type HTMLElementInterface interface {
	HTMlAttributes
	HTML_Renders
	HTML_Content
	HTML_ID_CLass_Styles
}

type HTML_Renders interface {
	RenderHTML() typesdef.HTML
}
type HTMlAttributes interface {
	AttributesOps
}

type HTML_ID_CLass_Styles interface {
	GetId() string
	GetClass() string
	GetStyle() string
}

type HTML_Content interface {
	SetInnerText(text string)
	GetInnerText() string
	GetChildren() []HTMLElementInterface
	AddChild(child HTMLElementInterface)
}
