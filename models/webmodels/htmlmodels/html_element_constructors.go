package htmlmodels

func NewHTMLElement(tagName string, attributes Attributes, children []HTMLElementInterface, innerText string) HTMLElement {
	return HTMLElement{tagName: tagName, attributes: attributes, children: children, innerText: innerText}
}
