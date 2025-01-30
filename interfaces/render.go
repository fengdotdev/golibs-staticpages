package interfaces

import "github.com/fengdotdev/golibs-staticpages/typesdef"

type Render interface {
	RenderHTML() typesdef.HTML
}

type WebRender interface {
	RenderHTML() typesdef.HTML
	RenderCSS
	RenderJS() typesdef.JS
}
