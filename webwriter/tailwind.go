package webwriter

import (
	"github.com/fengdotdev/golibs-staticpages/typesdef"
)

type Tailwind struct {
	companionType typesdef.CompanionType
	url           string
	position      typesdef.CompanionPosition
}

func NewTailwindCompanion() *Tailwind {
	return &Tailwind{
		position:      typesdef.Head,
		companionType: typesdef.RemoteJs,
		url:           "https://unpkg.com/@tailwindcss/browser@4",
	}
}

func (t *Tailwind) GetType() typesdef.CompanionType {
	return t.companionType
}

func (t *Tailwind) GetPosition() typesdef.CompanionPosition {
	return t.position
}

// return the html content of the companion
func (t *Tailwind) GetContent() string {
	return t.url
}

func (t *Tailwind) ForHTML() string {
	return "<script src=\"" + t.url + "\"></script>"
}

func (t *Tailwind) SetContent(content string) {
	t.url = content
}

func (t *Tailwind) IsValid() bool {
	return t.url != ""
}
