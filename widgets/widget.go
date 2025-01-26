package widgets

import (
	"github.com/fengdotdev/golibs-staticpages/interfaces"
	traits "github.com/fengdotdev/golibs-traits"
)

type Widget interface {
	interfaces.ElementHTML
	traits.JSONTrait
}
