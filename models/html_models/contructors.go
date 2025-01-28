package htmlmodels

func NewH1(attributes map[string]string, children []HTMLElement, innerText string) *HTMLElement {
	return &HTMLElement{
		TagName:    "h1",
		Attributes: attributes,
		Children:   children,
		InnerText:  innerText,
	}
}

func NewH2(attributes map[string]string, children []HTMLElement, innerText string) *HTMLElement {
	return &HTMLElement{
		TagName:    "h2",
		Attributes: attributes,
		Children:   children,
		InnerText:  innerText,
	}
}

func NewH3(attributes map[string]string, children []HTMLElement, innerText string) *HTMLElement {
	return &HTMLElement{
		TagName:    "h3",
		Attributes: attributes,
		Children:   children,
		InnerText:  innerText,
	}
}

func NewH4(attributes map[string]string, children []HTMLElement, innerText string) *HTMLElement {
	return &HTMLElement{
		TagName:    "h4",
		Attributes: attributes,
		Children:   children,
		InnerText:  innerText,
	}
}


