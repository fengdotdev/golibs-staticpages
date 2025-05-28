package htmlmodels

// constructors for Attributes
func NewAttributesEmpty() Attributes {
	return Attributes{attributes: make(map[string]string)}
}

func NewAttributes(attributes map[string]string) Attributes {
	return Attributes{attributes: attributes}
}

// Constructors for common attributes

func NewAttributesClass(class string) Attributes {
	return Attributes{attributes: map[string]string{"class": class}}
}

func NewAttributesID(id string) Attributes {
	return Attributes{attributes: map[string]string{"id": id}}
}

func NewAttributesStyle(style string) Attributes {
	return Attributes{attributes: map[string]string{"style": style}}
}

func NewAttributesClassAndID(class, id string) Attributes {
	return Attributes{attributes: map[string]string{"class": class, "id": id}}
}

func NewAttributesClassAndStyle(class, style string) Attributes {
	return Attributes{attributes: map[string]string{"class": class, "style": style}}
}

func NewAttributesIDAndStyle(id, style string) Attributes {
	return Attributes{attributes: map[string]string{"id": id, "style": style}}
}

func NewAttributesClassIDAndStyle(class, id, style string) Attributes {
	return Attributes{attributes: map[string]string{"class": class, "id": id, "style": style}}
}
