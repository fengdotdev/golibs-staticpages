package widgets_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/widgets"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestTextWidgetInterface(t *testing.T) {

	var text widgets.Widget = widgets.NewSimpleText("Hello World")

	assert.TrueWithMessage(t, text != nil, "Text should not be nil")

}

func TestTextToPlainHTML(t *testing.T) {

	text := widgets.NewSimpleText("Hello World")

	html := text.RenderHTML()

	result := html.GetHTML()

	assert.Equal(t, "<p>Hello World</p>", result)

	foo := "foo"
	tailwindcss := "text-red-500"
	textWithId := widgets.NewText("Hello World", widgets.TextOptions{Id: foo}, tailwindcss)

	htmlWithId := textWithId.RenderHTML()
	resultWithId := htmlWithId.GetHTML()

	assert.Equal(t, "<p id=\"foo\" class=\"text-red-500\">Hello World</p>", resultWithId)
}
