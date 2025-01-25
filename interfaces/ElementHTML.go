package interfaces

import (
	types "github.com/fengdotdev/golibs-staticpages/types"
)

type ElementHTML interface {
	RenderHTML() types.HTML
}
