package agnostic

import "github.com/fengdotdev/golibs-staticpages/widgets"

type Footer struct {
	child widgets.Widget
}

func NewFooter(child widgets.Widget) *Footer {
	return &Footer{
		child: child,
	}
}

func (f *Footer) GetChild() widgets.Widget {
	return f.child
}
