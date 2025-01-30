package cssmodels

import "github.com/fengdotdev/golibs-staticpages/typesdef"

type CSSModel struct {
	content map[CSSProperty]string
}

func NewCSSModel() *CSSModel {
	return &CSSModel{
		content: make(map[CSSProperty]string),
	}
}

func (c *CSSModel) AddCSSKeyword(keyword CSSProperty, value string) {
	c.content[keyword] = value
}

func (c *CSSModel) RenderPreCSS() string {
	separator := ": "
	var css string
	for k, v := range c.content {
		css += string(k) + separator + v + ";"
	}
	return css
}

func (c *CSSModel) RenderCSS(Selector string) typesdef.CSS {

	return RenderCSS(Selector, c.RenderPreCSS())
}

func RenderCSS(Selector string, precss string) typesdef.CSS {
	css := Selector + " {" + precss + "}"
	return *typesdef.NewCSS(css)
}
