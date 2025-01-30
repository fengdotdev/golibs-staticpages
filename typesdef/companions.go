package typesdef

type CompanionType string

const (
	JSCompanion  CompanionType = "jsCompanion"
	CSSCompanion CompanionType = "cssCompanion"
	RemoteJs     CompanionType = "remoteJs"
	RemoteCss    CompanionType = "remoteCss"
)

type CompanionPosition string

const (
	Head   CompanionPosition = "head"
	Body   CompanionPosition = "body"
	Footer CompanionPosition = "footer"
)

type Companion interface {
	GetType() CompanionType
	GetPosition() CompanionPosition
	GetContent() string
	SetContent(content string)
	ForHTML() string
	IsValid() bool
}
