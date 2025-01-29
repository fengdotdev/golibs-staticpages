package htmlmodels

import "github.com/fengdotdev/golibs-staticpages/typesdef"

// HTMlAttributes

func (e *HTMLElement) SetAttributes(attributes map[string]string) {
	e.attributes = NewAttributes(attributes)
}

func (e *HTMLElement) GetAttributes() map[string]string {
	return e.attributes.GetAttributes()
}

func (e *HTMLElement) GetAttribute(key string) string {
	return e.attributes.GetAttribute(key)
}

func (e *HTMLElement) GetAttributeOrError(key string) (string, error) {
	return e.attributes.GetAttributeOrError(key)
}

func (e *HTMLElement) AddAttribute(key, value string) {
	e.attributes.AddAttribute(key, value)
}

func (e *HTMLElement) RemoveAttribute(key string) {
	e.attributes.RemoveAttribute(key)
}

func (e *HTMLElement) ExistAttribute(key string) bool {
	return e.attributes.ExistAttribute(key)
}

func (e *HTMLElement) ClearAttributes() {
	e.attributes.ClearAttributes()
}

// HTML_Renders

func (e *HTMLElement) RenderHTML() typesdef.HTML {
	var output string

	hasChildren := len(e.children) > 0
	haveAttributes := len(e.attributes.GetAttributes()) > 0


	this is shit TODO 
	if hasChildren {
		childrenOutput := ""
		for _, child := range e.children {
			childrenOutput += child.
		}

		if haveAttributes {
			output = StartHTMLAttributes(e.tagName, e.attributes.GetAttributes()) + childrenOutput + e.innerText + EndHTML(e.tagName)
			return *typesdef.NewHTML(output)
		}

		output = StartHtml(e.tagName) + childrenOutput + e.innerText + EndHtml(e.tagName)
		return *typesdef.NewHTML(output)
	}

	if haveAttributes {
		output = StartHTMLAttributes(e.tagName, e.attributes.GetAttributes()) + e.innerText + EndHTML(e.tagName)
		return *typesdef.NewHTML(output)
	}

	output = StartHtml(e.tagName) + e.innerText + EndHtml(e.tagName)
	return *typesdef.NewHTML(output)
}

// HTML_Content

func (e *HTMLElement) SetInnerText(text string) {
	e.innerText = text
}

func (e *HTMLElement) GetInnerText() string {
	return e.innerText
}

func (e *HTMLElement) GetChildren() []HTMLElementInterface {
	return e.children
}

func (e *HTMLElement) AddChild(child HTMLElementInterface) {
	e.children = append(e.children, child)
}

// HTML_ID_CLass_Styles
func (e *HTMLElement) GetId() string {
	return e.attributes.GetID()
}

func (e *HTMLElement) GetClass() string {
	return e.attributes.GetClass()
}

func (e *HTMLElement) GetStyle() string {
	return e.attributes.GetStyle()
}
