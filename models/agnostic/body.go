package agnostic

import "github.com/fengdotdev/golibs-staticpages/widgets"

type Body struct {
	child widgets.Widget
}

func NewBody(child widgets.Widget) *Body {
	return &Body{
		child: child,
	}
}

func (b *Body) GetChild() widgets.Widget {
	return b.child
}
