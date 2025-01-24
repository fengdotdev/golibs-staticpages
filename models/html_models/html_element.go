package htmlmodels

type HTMLElement struct {
	TagName    string
	Attributes map[string]string
	Children   []HTMLElement
	InnerText  string
}
