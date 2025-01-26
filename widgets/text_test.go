package widgets_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/widgets"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestText(t *testing.T) {

	var text widgets.Widget = widgets.NewSimpleText("Hello World")

	assert.TrueWithMessage(t, text != nil, "Text should not be nil")

}
