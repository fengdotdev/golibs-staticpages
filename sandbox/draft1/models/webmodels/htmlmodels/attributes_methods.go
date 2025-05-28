package htmlmodels

import (
	"errors"
	"strings"
)

func (a *Attributes) GetAttributes() map[string]string {
	return a.attributes
}

func (a *Attributes) SetAttributes(attributes map[string]string) {
	a.attributes = attributes
}

// return a empty string if the key is not found
func (a *Attributes) GetAttribute(key string) string {
	if val, ok := a.attributes[key]; ok {
		return val
	}
	return ""

}

func (a *Attributes) GetAttributeOrError(key string) (string, error) {
	if val, ok := a.attributes[key]; ok {
		return val, nil
	}
	return "", errors.New("Attribute not found")
}

// add a new attribute to the map or update the value of an existing attribute
func (a *Attributes) AddAttribute(key, value string) {
	a.attributes[key] = value
}

func (a *Attributes) AppendAttribute(key, value string) {
	if val, ok := a.attributes[key]; ok {
		a.attributes[key] = val + value
		return
	}
	a.AddAttribute(key, value)
}

func (a *Attributes) RemoveAttribute(key string) {
	delete(a.attributes, key)
}

func (a *Attributes) ExistAttribute(key string) bool {
	_, ok := a.attributes[key]
	return ok
}

func (a *Attributes) ClearAttributes() {
	a.attributes = make(map[string]string)
}

// special types of attributes
func (a *Attributes) GetClass() string {
	return a.GetAttribute("class")
}

func (a *Attributes) SetClass(class string) {
	a.AddAttribute("class", class)
}

func (a *Attributes) GetID() string {
	return a.GetAttribute("id")
}

func (a *Attributes) SetID(id string) {
	a.AddAttribute("id", id)
}

func (a *Attributes) GetStyle() string {
	return a.GetAttribute("style")
}

func (a *Attributes) SetStyle(style string) {
	a.AddAttribute("style", style)
}

//AttributesMisc

func (a *Attributes) NumberOfAttributes() int {
	return len(a.attributes)
}

func (a *Attributes) ExistAttributes() bool {
	return len(a.attributes) > 0
}

// convertions

func (a *Attributes) ToMap() map[string]string {
	if !a.notNullOrEmpty() {
		return make(map[string]string)
	}
	return a.attributes
}

func (a *Attributes) ToSlice() []string {
	if !a.notNullOrEmpty() {
		return []string{}
	}
	var arr []string
	for key, value := range a.attributes {
		arr = append(arr, key+`="`+value+`"`)
	}
	return arr
}

func (a *Attributes) ToString() string {
	if !a.notNullOrEmpty() {
		return ""
	}
	var builder strings.Builder
	for key, value := range a.attributes {
		builder.WriteString(key + `="` + value + `" `)
	}
	result := builder.String()
	return strings.TrimSpace(result)
}

func (a *Attributes) notNullOrEmpty() bool {
	return a.attributes != nil && len(a.attributes) > 0 // but this is more readable bitch
}
