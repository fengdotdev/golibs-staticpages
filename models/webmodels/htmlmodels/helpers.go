package htmlmodels

// example: StartHtml("div") -> "<div>"
func StartHtml(tag string) string {
	return "<" + tag + ">"
}

// default func
func StartHTMLAttributes(tag string, attributes AttributesInterface) string {
	if attributes.ExistAttributes() {
		return "<" + tag + " " + attributes.ToString() + ">"
	}
	return "<" + tag + ">" 
}

// example: StartHtmlAttributes("div", map[string]string{"class": "text-red-500"}) -> "<div class='text-red-500'>"
func StartHTMLAttributesMap(tag string, attributes map[string]string) string {
	panic("StartHTMLAttributesMap redo this")
	if len(attributes) == 0 || attributes == nil {
		return StartHtml(tag)
	}

	var output string
	output = "<" + tag + " "
	for key, value := range attributes {
		output += key + "='" + value + "' "
	}
	output += ">"
	return output
}

// example: EndHtml("div") -> "</div>"
func EndHTML(tag string) string {
	return "</" + tag + ">"
}
