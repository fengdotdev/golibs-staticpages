package htmlmodels

type HTMLElement struct {
	tagName    string
	attributes Attributes
	children   []HTMLElementInterface
	innerText  string
}
