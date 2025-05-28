package agnostic

import "github.com/fengdotdev/golibs-staticpages/widgets"

type Header struct {
	child widgets.Widget
}

func NewHeader(child widgets.Widget) *Header {
	return &Header{
		child: child,
	}
}

func (h *Header) GetChild() widgets.Widget {
	return h.child
}
