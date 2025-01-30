package widgets

import (
	"github.com/fengdotdev/golibs-staticpages/interfaces"
	traits "github.com/fengdotdev/golibs-traits"
)

type Widget interface {
	interfaces.Render
	WidgetMisc
	interfaces.WidgetConversions
	traits.JSONTrait
}

type WidgetMisc interface {
	HaveChildren() bool
	NumbOfChildren() int
	HaveOptions() bool
}
