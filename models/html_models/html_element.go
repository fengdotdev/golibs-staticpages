package htmlmodels


type HTMLElement struct {
	tagName    string
	attributes Attributes
	children   []HTMLElement
	innerText  string
}

func (e *HTMLElement) GetHTMLString() string {
	return ""
}

func (e *HTMLElement) AddChild(child HTMLElement) {
	e.children = append(e.children, child)
}

func (e *HTMLElement) AddAttribute(key, value string) {
	e.attributes[key] = value
}

func (e *HTMLElement) SetInnerText(text string) {
	e.innerText = text
}

func (e *HTMLElement) GetInnerText() string {
	return e.innerText
}

func (e *HTMLElement) GetChildren() []HTMLElement {
	return e.children
}

func (e *HTMLElement) GetTagName() string {
	return e.tagName
}

func (e *HTMLElement) GetAttributes() map[string]string {
	return e.attributes
}

func (e *HTMLElement) GetAttribute(key string) string {
	return e.attributes[key]
}

func (e *HTMLElement) GetChild(index int) HTMLElement {
	return e.children[index]
}

func (e *HTMLElement) GetChildCount() int {
	return len(e.children)
}
