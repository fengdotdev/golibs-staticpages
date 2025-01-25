package interfaces

import (
	"github.com/fengdotdev/golibs-staticpages/typesdef"
)

type ElementHTML interface {
	RenderHTML() typesdef.HTML
}
