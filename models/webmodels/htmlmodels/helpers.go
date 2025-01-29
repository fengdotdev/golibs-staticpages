package htmlmodels

// example: StartHtml("div") -> "<div>"
func StartHtml(tag string) string {
	return "<" + tag + ">"
}

// example: StartHtmlWith("div", "class='text-red-500'") -> "<div class='text-red-500'>"
func StartHTMLWith(tag string, with string) string {
	return "<" + tag + " " + with + ">"
}

// example: StartHtmlAttributes("div", map[string]string{"class": "text-red-500"}) -> "<div class='text-red-500'>"
// prefer this function over StartHtmlWith and use a map of attributes instead of a string
// if u dont have any attributes, pass nil
func StartHTMLAttributes(tag string, attributes map[string]string) string {
zsasadadasas
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
