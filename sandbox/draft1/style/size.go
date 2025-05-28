package style

import (
	"fmt"

	"github.com/fengdotdev/golibs-staticpages/models/webmodels/cssmodels"
)

type Size struct {
	Width  int
	Height int
}

func (s *Size) WidthString() string {
	return fmt.Sprintf("%d", s.Width)
}

func (s *Size) HeightString() string {
	return fmt.Sprintf("%d", s.Height)
}

func (s *Size) RenderPreCSS() string {

	css := cssmodels.NewCSSModel()
	css.AddCSSKeyword(cssmodels.Width, s.WidthString())
	css.AddCSSKeyword(cssmodels.Height, s.HeightString())

	return css.RenderPreCSS()
}
