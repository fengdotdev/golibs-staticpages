package style

import "github.com/fengdotdev/golibs-staticpages/typesdef"

type Style struct {
	Size            Size
	Font            Font
	Border          Border
	Padding         Padding
	Margin          Margin
	BackgroundColor Color
	ForegroundColor Color
}

func NewStyle() *Style {
	return &Style{
		Size:            Size{},
		Font:            Font{},
		Border:          Border{},
		Padding:         Padding{},
		Margin:          Margin{},
		BackgroundColor: Color{},
		ForegroundColor: Color{},
	}
}

// css

func (s *Style) RenderCSS() typesdef.CSS {

	s.Size.RenderPreCSS()
	return typesdef.CSS{}

}
