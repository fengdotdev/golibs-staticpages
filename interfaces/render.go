package interfaces

import "github.com/fengdotdev/golibs-staticpages/typesdef"

type Render interface {
	RenderHTML() typesdef.HTML
}
