package interfaces

import (
	types "github.com/fengdotdev/golibs-staticpages/types"
)

type Element interface {
	RenderHTML() types.HTML
}
