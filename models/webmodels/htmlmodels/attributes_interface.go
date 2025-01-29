package htmlmodels

type AttributesInterface interface {
	AttributesOps
	AttributesConvertions
	AttributesID_CLass_Style
}

type AttributesOps interface {
	SetAttributes(attributes map[string]string)
	GetAttributes() map[string]string

	GetAttribute(key string) string
	GetAttributeOrError(key string) (string, error)
	AddAttribute(key, value string)

	RemoveAttribute(key string)
	ExistAttribute(key string) bool
	ClearAttributes()
}

type AttributesConvertions interface {
	ToString() string
	ToMap() map[string]string
	ToSlice() []string
}

type AttributesID_CLass_Style interface {
	GetClass() string
	SetClass(class string)
	GetID() string
	SetID(id string)
	GetStyle() string
	SetStyle(style string)
}
