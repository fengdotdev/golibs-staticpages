package agnostic

import "github.com/fengdotdev/golibs-staticpages/widgets"

type Agnostic interface {
	GetChild() widgets.Widget
}
