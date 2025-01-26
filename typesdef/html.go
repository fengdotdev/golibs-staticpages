package typesdef

import (
	"github.com/fengdotdev/golibs-staticpages/validators"
)

type HTML struct {
	content string
}

func NewHTML(content string) *HTML {
	return &HTML{content: content}
}

func (h *HTML) GetType() CompanionType {
	panic("Not implemented")
}

func (h *HTML) IsValid() bool {

	html := h.GetContent()
	return validators.HTMLisValid(html)
}

func (h *HTML) GetContent() string {
	return h.content
}

func (h *HTML) GetHTML() string {
	return h.GetContent()
}

func (h *HTML) SetContent(content string) {
	h.content = content
}
