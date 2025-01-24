package types

type CompanionType string

const (
	JSCompanion  CompanionType = "jsCompanion"
	CSSCompanion CompanionType = "cssCompanion"
)

type Companion interface {
	GetType() CompanionType
	GetContent() string
	SetContent(content string)
	IsValid() bool
}
