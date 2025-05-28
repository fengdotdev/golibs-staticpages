package typesdef

import "github.com/fengdotdev/golibs-staticpages/validators"

type CSS struct {
	content string
}

func NewCSS(content string) *CSS {
	return &CSS{
		content: content,
	}
}

func (c *CSS) GetType() CompanionType {
	return CSSCompanion
}

func (c *CSS) GetContent() string {
	return c.content
}

func (c *CSS) SetContent(content string) {
	c.content = content
}

func (c *CSS) IsValid() bool {
	return validators.CSSIsValid(c.content)
}
