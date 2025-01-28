package htmlmodels

type HTMLElementInterface interface {
	GetHTMLString() string
	AddChild(child HTMLElementInterface)
	AddAttribute(key, value string)
	SetInnerText(text string)
	GetInnerText() string
	GetChildren() []HTMLElementInterface
}


